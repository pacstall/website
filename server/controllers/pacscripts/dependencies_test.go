package psapi

import (
	"testing"

	"pacstall.dev/webserver/config"
	pkgcache "pacstall.dev/webserver/services/package_cache"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/utils/expect"
)

func Test_PackageController_findPackageDependencies_NotFound(t *testing.T) {
	cacheService := pkgcache.New()
	controller := New(config.ServerConfiguration{}, cacheService)

	_, found := controller.findPackageDependencies("package-that-does-not-exist")
	expect.False(t, "expected to not find any package", found)
}

func Test_PackageController_findPackageDependencies_Found(t *testing.T) {
	cacheService := pkgcache.New()
	controller := New(config.ServerConfiguration{}, cacheService)

	cacheService.Update([]*pac.Script{
		{
			PackageName: "test",
		},
	})

	dependencies, found := controller.findPackageDependencies("test")
	expect.True(t, "expected to find test package", found)

	expected := pacscriptDependencies{
		PacstallDependencies: []*pac.Script{},
	}

	expect.Equals(t, "dependencies", expected, dependencies)
}

func Test_PackageController_findPackageDependencies_CorrectDependencies(t *testing.T) {
	cacheService := pkgcache.New()
	controller := New(config.ServerConfiguration{}, cacheService)

	cacheService.Update([]*pac.Script{
		{
			PackageName:          "test",
			RuntimeDependencies:  []string{"runtime"},
			BuildDependencies:    []string{"build"},
			OptionalDependencies: []string{"optional"},
			PacstallDependencies: []string{"test-dependency"},
		},
		{
			PackageName: "test-dependency",
		},
	})

	dependencies, found := controller.findPackageDependencies("test")
	expect.True(t, "expected to find test package", found)

	expected := pacscriptDependencies{
		RuntimeDependencies:  []string{"runtime"},
		BuildDependencies:    []string{"build"},
		OptionalDependencies: []string{"optional"},
		PacstallDependencies: []*pac.Script{
			{
				PackageName: "test-dependency",
			},
		},
	}

	expect.Equals(t, "dependencies", expected, dependencies)
}
