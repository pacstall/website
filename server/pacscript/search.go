package pacscript

import (
	"sort"
	"strings"

	"pacstall.dev/website/types"
)

const DEFAULT = "default"

var SortableProperties = []string{"maintainer", "version", "name", DEFAULT}
var SortTypes = []string{"asc", "desc", DEFAULT}
var FilterableProperties = []string{"name", "maintainer", DEFAULT}

const (
	PageKey     = "page"
	SizeKey     = "size"
	SortByKey   = "sortBy"
	SortKey     = "sort"
	FilterByKey = "filterBy"
	FilterKey   = "filter"
)

func FindPackageInList(name string, packages []*types.PackageInfo) *types.PackageInfo {
	for _, pkg := range packages {
		if strings.Compare(pkg.Name, name) == 0 {
			return pkg
		}
	}

	return nil
}

func FilterPackages(packages []*types.PackageInfo, filter, filterBy string) []*types.PackageInfo {
	filterByFunc := func(matches func(*types.PackageInfo) bool) []*types.PackageInfo {
		out := make([]*types.PackageInfo, 0)
		for _, pkg := range packages {
			if matches(pkg) {
				out = append(out, pkg)
			}
		}
		return out
	}

	switch filterBy {
	case "name":
		return filterByFunc(func(pi *types.PackageInfo) bool {
			return strings.Contains(pi.Name, filter) ||
				strings.Contains(pi.PackageName, filter) ||
				strings.Contains(pi.Gives, filter) ||
				strings.Contains(pi.Description, filter)
		})

	case "maintainer":
		return filterByFunc(func(pi *types.PackageInfo) bool {
			return strings.Contains(pi.Maintainer, filter)
		})
	default:
		return packages
	}
}

func SortPackages(packages []*types.PackageInfo, sortType, sortBy string) []*types.PackageInfo {
	if sortType == DEFAULT {
		return packages
	}

	clone := make([]*types.PackageInfo, 0)
	clone = append(clone, packages...)

	switch sortBy {
	case "name":
		if strings.Compare(sortType, "asc") == 0 {
			sort.Sort(SortByName{clone})
		} else {
			sort.Sort(SortByNameDesc{clone})
		}

	case "maintainer":
		if strings.Compare(sortType, "asc") == 0 {
			sort.Sort(SortByMaintainer{clone})
		} else {
			sort.Sort(SortByMaintainerDesc{clone})
		}

	case "version":
		if strings.Compare(sortType, "asc") == 0 {
			sort.Sort(SortByVersion{clone})
		} else {
			sort.Sort(SortByVersionDesc{clone})
		}
	default:
		break
	}

	return clone
}
