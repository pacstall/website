package repology_api

import (
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/types/controller"
	"pacstall.dev/webserver/types/service"
)

type RepologyController struct {
	serverConfiguration config.ServerConfiguration
	packageCacheService service.PackageCacheService
}

func New(serverConfiguration config.ServerConfiguration, packageCacheService service.PackageCacheService) *RepologyController {
	c := &RepologyController{}

	c.serverConfiguration = serverConfiguration
	c.packageCacheService = packageCacheService

	return c
}

func (c *RepologyController) GetRoutes() []controller.ControllerRoute {
	return []controller.ControllerRoute{
		{
			Method: controller.METHOD_GET,
			Path:   "/api/repology",
			Handle: c.GetRepologyPackageListHandle,
		},
	}
}
