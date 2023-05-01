package urlshortener

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"pacstall.dev/webserver/model"
)

var db = model.Instance()
var incrementVisitsExpression = gorm.Expr(model.ShortenedLinkColumns.Visits+" + ?", 1)

func GetShortenedLinkRedirectHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	linkId := params["linkId"]

	if linkId == "" {
		w.WriteHeader(404)
		return
	}

	var shortenedLink model.ShortenedLink
	if result := db.Where(model.ShortenedLink{LinkId: params["linkId"]}).First(&shortenedLink); result.Error != nil {
		w.WriteHeader(404)
		return
	}

	db.Model(&shortenedLink).Update(model.ShortenedLinkColumns.Visits, incrementVisitsExpression)

	w.Header().Add("Location", shortenedLink.Link)
	w.WriteHeader(302)
}
