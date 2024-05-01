package service

type MatomoTrackerService interface {
	TrackShortLink(user, userAgent, urlRef, link string)
}
