package matomo_tracker

import (
	"strings"

	"github.com/treelightsoftware/go-matomo"
	"pacstall.dev/webserver/log"
)

type MatomoTrackerService struct{}

func New() *MatomoTrackerService {
	return &MatomoTrackerService{}
}

func (s *MatomoTrackerService) TrackShortLink(user, userAgent, urlRef, link string) {
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
		log.Warn("failed to ping matomo tracker: %+v", err)
	}
}
