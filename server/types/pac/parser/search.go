package parser

import (
	"strings"

	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
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

func FilterPackages(packages []*pac.Script, filter, filterBy string) []*pac.Script {
	filterByFunc := func(matches func(*pac.Script) bool) []*pac.Script {
		out := make([]*pac.Script, 0)
		for _, pkg := range packages {
			if matches(pkg) {
				out = append(out, pkg)
			}
		}
		return out
	}

	switch filterBy {
	case "name":
		return filterByFunc(func(pi *pac.Script) bool {
			return strings.Contains(pi.PackageName, filter) ||
				strings.Contains(pi.Gives, filter) ||
				strings.Contains(pi.Description, filter)
		})

	case "maintainer":
		return filterByFunc(func(pi *pac.Script) bool {
			return strings.Contains(strings.Join(pi.Maintainers, ", "), filter)
		})
	default:
		return packages
	}
}

func SortPackages(packages []*pac.Script, sortType, sortBy string) []*pac.Script {
	if sortType == DEFAULT {
		return packages
	}

	out := array.Clone(packages)

	switch sortBy {
	case "name":
		if strings.Compare(sortType, "asc") == 0 {
			out = array.SortBy(out, func(a, b *pac.Script) bool {
				return strings.Compare(a.PackageName, b.PackageName) < 0
			})
		} else {
			out = array.SortBy(out, func(a, b *pac.Script) bool {
				return strings.Compare(a.PackageName, b.PackageName) > 0
			})
		}

	case "maintainer":
		if strings.Compare(sortType, "asc") == 0 {
			out = array.SortBy(out, func(a, b *pac.Script) bool {
				return strings.Compare(strings.Join(a.Maintainers, ","), strings.Join(b.Maintainers, ",")) < 0
			})
		} else {
			out = array.SortBy(out, func(a, b *pac.Script) bool {
				return strings.Compare(strings.Join(a.Maintainers, ","), strings.Join(b.Maintainers, ",")) > 0
			})
		}

	case "version":
		if strings.Compare(sortType, "asc") == 0 {
			out = array.SortBy(out, func(a, b *pac.Script) bool {
				return strings.Compare(a.Version, b.Version) < 0
			})
		} else {
			out = array.SortBy(out, func(a, b *pac.Script) bool {
				return strings.Compare(a.Version, b.Version) > 0
			})
		}
	default:
		break
	}

	return out
}
