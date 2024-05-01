package psapi

import (
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/types/controller"
	"pacstall.dev/webserver/types/service"
)

type PackageController struct {
	serverConfiguration config.ServerConfiguration
	packageCacheService service.PackageCacheService
}

func New(serverConfiguration config.ServerConfiguration, packageCacheService service.PackageCacheService) *PackageController {
	c := &PackageController{}

	c.serverConfiguration = serverConfiguration
	c.packageCacheService = packageCacheService

	return c
}

func (c *PackageController) GetRoutes() []controller.ControllerRoute {
	return []controller.ControllerRoute{
		{
			Method: controller.METHOD_GET,
			Path:   "/api/packages",
			Handle: c.GetPackageListHandle,
		},
		{
			Method: controller.METHOD_GET,
			Path:   "/api/packages/{name}",
			Handle: c.GetPackageHandle,
		},
		{
			Method: controller.METHOD_GET,
			Path:   "/api/packages/{name}/dependencies",
			Handle: c.GetPackageDependenciesHandle,
		},
		{
			Method: controller.METHOD_GET,
			Path:   "/api/packages/{name}/requiredBy",
			Handle: c.GetPackageRequiredByHandle,
		},
	}
}
