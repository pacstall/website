package psapi

import (
	"testing"

	"pacstall.dev/webserver/config"
	pkgcache "pacstall.dev/webserver/services/package_cache"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/utils/expect"
)

func Test_PackageController_findPackageByName_NotFound(t *testing.T) {
	cacheService := pkgcache.New()
	controller := New(config.ServerConfiguration{}, cacheService)

	_, found := controller.findPackageByName("package-that-does-not-exist")

	expect.False(t, "expected to not find any package", found)
}

func Test_PackageController_findPackageByName_Found(t *testing.T) {
	cacheService := pkgcache.New()
	controller := New(config.ServerConfiguration{}, cacheService)

	cacheService.Update([]*pac.Script{
		{
			PackageName: "test",
		},
	})

	dependencies, found := controller.findPackageByName("test")
	expect.True(t, "expected to find test package", found)

	expected := &pac.Script{
		PackageName: "test",
	}

	expect.Equals(t, "dependencies", expected, dependencies)
}
