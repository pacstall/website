package urlshortener

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/treelightsoftware/go-matomo"
	"gorm.io/gorm"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/model"
)

var db = model.Instance()
var incrementVisitsExpression = gorm.Expr(model.ShortenedLinkColumns.Visits+" + ?", 1)

var pathParams = struct {
	LinkId string
}{
	LinkId: "linkId",
}

var queryParams = struct {
	DoNotTrack string
}{
	DoNotTrack: "dnt",
}

func GetShortenedLinkRedirectHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	query := req.URL.Query()

	linkId := params[pathParams.LinkId]
	_, doNotTrack := query[queryParams.DoNotTrack]

	if linkId == "" {
		w.WriteHeader(404)
		return
	}

	var shortenedLink model.ShortenedLink
	if result := db.Where(model.ShortenedLink{LinkId: linkId}).First(&shortenedLink); result.Error != nil {
		w.WriteHeader(404)
		return
	}

	// Increment visits in the background and ping matomo
	go func() {
		if doNotTrack {
			return
		}

		db.Model(&shortenedLink).Update(model.ShortenedLinkColumns.Visits, incrementVisitsExpression)
		if config.Matomo.Enabled {
			pingMatomoTracker(req.RemoteAddr, req.UserAgent(), req.Referer(), linkId)
		}
	}()

	w.Header().Add("Location", shortenedLink.Link)
	w.WriteHeader(302)
}

func pingMatomoTracker(user, userAgent, urlRef, link string) {
	// Strip the port from the user
	user = user[:strings.LastIndex(user, ":")]

	params := &matomo.Parameters{
		RecommendedParameters: &matomo.RecommendedParameters{
			URL:        matomo.StringPtr("/" + link),
			ActionName: matomo.StringPtr("ShortenedLink"),
			VisitorID:  matomo.StringPtr("@pacstall/webserver/" + user),
		},
		UserParameters: &matomo.UserParameters{
			UserID:    matomo.StringPtr("@pacstall/webserver/" + user),
			UserAgent: matomo.StringPtr(userAgent),
			URLRef:    matomo.StringPtr(urlRef),
		},
		EventTrackingParameters: &matomo.EventTrackingParameters{
			Category: matomo.StringPtr("ShortenedLink"),
			Action:   matomo.StringPtr(link),
		},
	}

	err := matomo.Send(params)
	if err != nil {
		log.Warn("failed to ping matomo tracker: %s", err)
	}
}
