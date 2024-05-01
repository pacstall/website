package psapi

import (
	"testing"

	"pacstall.dev/webserver/config"
	pkgcache "pacstall.dev/webserver/services/package_cache"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/utils/expect"
)

func Test_PackageController_findPackagesRequiredBy_NotFound(t *testing.T) {
	cacheService := pkgcache.New()
	controller := New(config.ServerConfiguration{}, cacheService)

	_, found := controller.findPackagesRequiredBy("package-that-does-not-exist")
	expect.False(t, "expected to not find any package", found)
}

func Test_PackageController_findPackagesRequiredBy_Found(t *testing.T) {
	cacheService := pkgcache.New()
	controller := New(config.ServerConfiguration{}, cacheService)

	cacheService.Update([]*pac.Script{
		{
			PackageName: "test",
			RequiredBy: []string{
				"package_1",
				"package_2",
			},
		},
		{
			PackageName: "package_1",
			PacstallDependencies: []string{
				"test",
			},
		},
		{
			PackageName: "package_2",
			PacstallDependencies: []string{
				"test",
			},
		},
	})

	requiredBy, found := controller.findPackagesRequiredBy("test")
	expect.True(t, "expected to find test package", found)

	expected := []*pac.Script{
		{
			PackageName: "package_1",
			PacstallDependencies: []string{
				"test",
			},
		},
		{
			PackageName: "package_2",
			PacstallDependencies: []string{
				"test",
			},
		},
	}

	expect.Equals(t, "dependencies", expected, requiredBy)
}
