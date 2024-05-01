package psapi

import (
	"testing"

	"pacstall.dev/webserver/config"
	pkgcache "pacstall.dev/webserver/services/package_cache"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/utils/expect"
)

var testData = []*pac.Script{
	{
		PackageName: "test_1",
		Version:     "1.0.0",
		Maintainers: []string{
			"Paul Cosma <paul.cosma97@gmail.com>",
			"saenai255 <paul.cosma97@gmail.com>",
		},
	},
	{
		PackageName: "test_2",
		Version:     "2.3.1",
		Maintainers: []string{
			"John Doe",
		},
	},
	{
		PackageName: "test_3",
		Version:     "2.3.1",
		Maintainers: []string{
			"Jane Doe",
		},
	},
}

func Test_PackageController_findProjectsPageMatchingFilter_Empty(t *testing.T) {
	cacheService := pkgcache.New()
	controller := New(config.ServerConfiguration{}, cacheService)

	filter := ""
	filterBy := ""
	sort := ""
	sortBy := ""
	page := 0
	pageSize := 50

	results := controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected := packageListPage{
		Page:     0,
		Size:     50,
		LastPage: 0,
		Sort:     "",
		SortBy:   "",
		Filter:   "",
		FilterBy: "",
		Total:    0,
		Data:     []*pac.Script{},
	}

	expect.Equals(t, "package page", expected, results)
}

func Test_PackageController_findProjectsPageMatchingFilter_SomeResults(t *testing.T) {
	cacheService := pkgcache.New()
	cacheService.Update(testData)
	controller := New(config.ServerConfiguration{}, cacheService)

	filter := ""
	filterBy := ""
	sort := ""
	sortBy := ""
	page := 0
	pageSize := 5

	results := controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected := packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     "",
		SortBy:   "",
		Filter:   "",
		FilterBy: "",
		Total:    3,
		Data:     testData,
	}

	expect.Equals(t, "package page", expected, results)
}

func Test_PackageController_findProjectsPageMatchingFilter_PackageNameFilter(t *testing.T) {
	cacheService := pkgcache.New()
	cacheService.Update(testData)
	controller := New(config.ServerConfiguration{}, cacheService)

	filter := "test"
	filterBy := "name"
	sort := ""
	sortBy := ""
	page := 0
	pageSize := 5

	results := controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected := packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    3,
		Data: []*pac.Script{
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
			{
				PackageName: "test_3",
				Version:     "2.3.1",
				Maintainers: []string{
					"Jane Doe",
				},
			},
		},
	}

	expect.Equals(t, "package page", expected, results)

	filter = "1"
	results = controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected = packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    1,
		Data: []*pac.Script{
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
		},
	}

	expect.Equals(t, "package page", expected, results)
}

func Test_PackageController_findProjectsPageMatchingFilter_MaintainerFilter(t *testing.T) {
	cacheService := pkgcache.New()
	cacheService.Update(testData)
	controller := New(config.ServerConfiguration{}, cacheService)

	filter := "saenai255"
	filterBy := "maintainer"
	sort := ""
	sortBy := ""
	page := 0
	pageSize := 5

	results := controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected := packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    1,
		Data: []*pac.Script{
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
		},
	}

	expect.Equals(t, "package page", expected, results)

	filter = "doe"
	results = controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected = packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    2,
		Data: []*pac.Script{

			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
			{
				PackageName: "test_3",
				Version:     "2.3.1",
				Maintainers: []string{
					"Jane Doe",
				},
			},
		},
	}

	expect.Equals(t, "package page", expected, results)
}

func Test_PackageController_findProjectsPageMatchingFilter_SortName(t *testing.T) {
	cacheService := pkgcache.New()
	cacheService.Update(testData)
	controller := New(config.ServerConfiguration{}, cacheService)

	filter := ""
	filterBy := ""
	sort := "asc"
	sortBy := "name"
	page := 0
	pageSize := 5

	results := controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected := packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    3,
		Data: []*pac.Script{
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
			{
				PackageName: "test_3",
				Version:     "2.3.1",
				Maintainers: []string{
					"Jane Doe",
				},
			},
		},
	}

	expect.Equals(t, "package page sorted ascending", expected, results)

	sort = "desc"
	results = controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected = packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    3,
		Data: []*pac.Script{
			{
				PackageName: "test_3",
				Version:     "2.3.1",
				Maintainers: []string{
					"Jane Doe",
				},
			},
			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
		},
	}

	expect.Equals(t, "package page sorted descending", expected, results)
}

func Test_PackageController_findProjectsPageMatchingFilter_SortVersion(t *testing.T) {
	cacheService := pkgcache.New()
	cacheService.Update(testData)
	controller := New(config.ServerConfiguration{}, cacheService)

	filter := ""
	filterBy := ""
	sort := "asc"
	sortBy := "name"
	page := 0
	pageSize := 5

	results := controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected := packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    3,
		Data: []*pac.Script{
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
			{
				PackageName: "test_3",
				Version:     "2.3.1",
				Maintainers: []string{
					"Jane Doe",
				},
			},
		},
	}

	expect.Equals(t, "package page sorted ascending", expected, results)

	sort = "desc"
	results = controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected = packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    3,
		Data: []*pac.Script{
			{
				PackageName: "test_3",
				Version:     "2.3.1",
				Maintainers: []string{
					"Jane Doe",
				},
			},
			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
		},
	}

	expect.Equals(t, "package page sorted descending", expected, results)
}

func Test_PackageController_findProjectsPageMatchingFilter_SortMaintainer(t *testing.T) {
	cacheService := pkgcache.New()
	cacheService.Update(testData)
	controller := New(config.ServerConfiguration{}, cacheService)

	filter := ""
	filterBy := ""
	sort := "asc"
	sortBy := "maintainer"
	page := 0
	pageSize := 5

	results := controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected := packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    3,
		Data: []*pac.Script{
			{
				PackageName: "test_3",
				Version:     "2.3.1",
				Maintainers: []string{
					"Jane Doe",
				},
			},
			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
		},
	}

	expect.Equals(t, "package page sorted ascending", expected, results)

	sort = "desc"
	results = controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected = packageListPage{
		Page:     0,
		Size:     5,
		LastPage: 0,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    3,
		Data: []*pac.Script{
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
			{
				PackageName: "test_3",
				Version:     "2.3.1",
				Maintainers: []string{
					"Jane Doe",
				},
			},
		},
	}

	expect.Equals(t, "package page sorted descending", expected, results)
}

func Test_PackageController_findProjectsPageMatchingFilter_Pagination(t *testing.T) {
	cacheService := pkgcache.New()
	cacheService.Update(testData)
	controller := New(config.ServerConfiguration{}, cacheService)

	filter := ""
	filterBy := ""
	sort := ""
	sortBy := ""
	page := 0
	pageSize := 2

	results := controller.findProjectsPageMatchingFilter(
		filter,
		filterBy,
		sort,
		sortBy,
		page,
		pageSize,
	)

	expected := packageListPage{
		Page:     0,
		Size:     2,
		LastPage: 1,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    3,
		Data: []*pac.Script{
			{
				PackageName: "test_1",
				Version:     "1.0.0",
				Maintainers: []string{
					"Paul Cosma <paul.cosma97@gmail.com>",
					"saenai255 <paul.cosma97@gmail.com>",
				},
			},
			{
				PackageName: "test_2",
				Version:     "2.3.1",
				Maintainers: []string{
					"John Doe",
				},
			},
		},
	}

	expect.Equals(t, "package page sorted ascending", expected, results)
}
