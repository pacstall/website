package urlshortener

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/treelightsoftware/go-matomo"
	"pacstall.dev/webserver/log"
)

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

func (c *UrlShortenerController) GetShortenedLinkRedirectHandle(w http.ResponseWriter, req *http.Request) error {
	params := mux.Vars(req)
	query := req.URL.Query()

	linkId := params[pathParams.LinkId]
	_, doNotTrack := query[queryParams.DoNotTrack]

	if linkId == "" {
		w.WriteHeader(404)
		return nil
	}

	shortenedLink, err := c.shortenedLinkRepository.FindOneByLinkId(linkId)
	if err != nil {
		w.WriteHeader(404)
		return nil
	}

	// Increment visits in the background and ping matomo
	go func() {
		if doNotTrack {
			return
		}

		if err := c.shortenedLinkRepository.IncrementVisits(shortenedLink.GetID()); err != nil {
			log.Warn("failed to increment visits. err: %+v", err)
		}

		if c.matomoConfiguration.Enabled {
			pingMatomoTracker(req.RemoteAddr, req.UserAgent(), req.Referer(), linkId)
		}
	}()

	w.Header().Add("Location", shortenedLink.GetLink())
	w.WriteHeader(302)

	return nil
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
