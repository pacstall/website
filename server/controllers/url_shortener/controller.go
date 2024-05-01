package urlshortener

import (
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/types/controller"
	"pacstall.dev/webserver/types/repository"
	"pacstall.dev/webserver/types/service"
)

type UrlShortenerController struct {
	matomoConfiguration     config.MatomoConfiguration
	shortenedLinkRepository repository.ShortenedLinkRepository
	matomoTrackerSerice     service.MatomoTrackerService
}

func New(
	matomoConfiguration config.MatomoConfiguration,
	shortenedLinkRepository repository.ShortenedLinkRepository,
	matomoTrackerSerice service.MatomoTrackerService,
) *UrlShortenerController {
	c := &UrlShortenerController{}

	c.matomoConfiguration = matomoConfiguration
	c.matomoTrackerSerice = matomoTrackerSerice
	c.shortenedLinkRepository = shortenedLinkRepository

	return c
}

func (c *UrlShortenerController) GetRoutes() []controller.ControllerRoute {
	return []controller.ControllerRoute{
		{
			Method: controller.METHOD_GET,
			Path:   "/q/{linkId}",
			Handle: c.GetShortenedLinkRedirectHandle,
		},
	}
}
