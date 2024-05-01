package urlshortener

import (
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/types/controller"
	"pacstall.dev/webserver/types/repository"
)

type UrlShortenerController struct {
	matomoConfiguration     config.MatomoConfiguration
	shortenedLinkRepository repository.ShortenedLinkRepository
}

func New(matomoConfiguration config.MatomoConfiguration) *UrlShortenerController {
	c := &UrlShortenerController{}

	c.matomoConfiguration = matomoConfiguration
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
