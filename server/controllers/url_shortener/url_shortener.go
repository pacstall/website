package urlshortener

import (
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/repository"
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

	shortenedLink, found := c.findShortenedLinkAndTrack(linkId, doNotTrack, req.RemoteAddr, req.UserAgent(), req.Referer())
	if !found {
		w.WriteHeader(404)
		return nil
	}

	w.Header().Add("Location", shortenedLink.GetLink())
	w.WriteHeader(302)

	return nil
}

func (c *UrlShortenerController) findShortenedLinkAndTrack(linkId string, doNotTrack bool, remoteAddress, userAgent, referer string) (repository.ShortenedLink, bool) {
	shortenedLink, err := c.shortenedLinkRepository.FindOneByLinkId(linkId)
	if err != nil {
		return nil, false
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
			c.matomoTrackerSerice.TrackShortLink(remoteAddress, userAgent, referer, linkId)
		}
	}()

	return shortenedLink, true
}
