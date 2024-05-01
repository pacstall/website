package repology_api

import (
	"testing"

	"pacstall.dev/webserver/config"
	pkgcache "pacstall.dev/webserver/services/package_cache"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/utils/expect"
)

func Test_RepologyController_findPackagesForRepology_Empty(t *testing.T) {
	cacheService := pkgcache.New()
	repologyService := New(
		config.ServerConfiguration{},
		cacheService,
	)

	actual := repologyService.findPackagesForRepology()
	expected := []repologyPackage{}

	expect.Equals(t, "repology packages", expected, actual)
}

func Test_RepologyController_findPackagesForRepology_WithPackages(t *testing.T) {
	cacheService := pkgcache.New()
	repologyService := New(
		config.ServerConfiguration{},
		cacheService,
	)

	cacheService.Update([]*pac.Script{
		{
			PackageName: "test-package-bin",
			Version:     "2.1.3",
			Maintainers: []string{"Paul Cosma <paul.cosma97@gmail.com>"},
		},
		{
			PackageName: "test-package",
			Version:     "1.0.0",
			Maintainers: []string{"Paul Cosma <paul.cosma97@gmail.com>"},
		},
		{
			PackageName: "test-package-git",
			Version:     "1.0.0",
			Maintainers: []string{"Paul Cosma <paul.cosma97@gmail.com>"},
		},
		{
			PackageName: "test-package-deb",
			Version:     "1.0.0",
			Maintainers: []string{"Paul Cosma <paul.cosma97@gmail.com>"},
		},
	})

	actual := repologyService.findPackagesForRepology()

	maintainerName := "Paul Cosma"
	maintainerEmail := "paul.cosma97@gmail.com"
	expected := []repologyPackage{
		{
			Name: "test-package-bin",
			Maintainer: maintainerDetails{
				Name:  &maintainerName,
				Email: &maintainerEmail,
			},
			Version:           "2.1.3",
			RecipeURL:         "https://raw.githubusercontent.com/pacstall/pacstall-programs/master/packages/test-package-bin/test-package-bin.pacscript",
			PackageDetailsURL: "https://pacstall.dev/packages/test-package-bin",
			Type:              "Precompiled",
		},
		{
			Name: "test-package",
			Maintainer: maintainerDetails{
				Name:  &maintainerName,
				Email: &maintainerEmail,
			},
			Version:           "1.0.0",
			RecipeURL:         "https://raw.githubusercontent.com/pacstall/pacstall-programs/master/packages/test-package/test-package.pacscript",
			PackageDetailsURL: "https://pacstall.dev/packages/test-package",
			Type:              "Source Code",
		},
		{
			Name: "test-package-git",
			Maintainer: maintainerDetails{
				Name:  &maintainerName,
				Email: &maintainerEmail,
			},
			Version:           "1.0.0",
			RecipeURL:         "https://raw.githubusercontent.com/pacstall/pacstall-programs/master/packages/test-package-git/test-package-git.pacscript",
			PackageDetailsURL: "https://pacstall.dev/packages/test-package-git",
			Type:              "Source Code",
		},
		{
			Name: "test-package-deb",
			Maintainer: maintainerDetails{
				Name:  &maintainerName,
				Email: &maintainerEmail,
			},
			Version:           "1.0.0",
			RecipeURL:         "https://raw.githubusercontent.com/pacstall/pacstall-programs/master/packages/test-package-deb/test-package-deb.pacscript",
			PackageDetailsURL: "https://pacstall.dev/packages/test-package-deb",
			Type:              "Debian Native",
		},
	}

	expect.Equals(t, "repology packages", expected, actual)
}
