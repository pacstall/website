package pacpkgs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"pacstall.dev/website/serverlib"
	"pacstall.dev/website/serverlib/query"
	"pacstall.dev/website/types"
)

const DEFAULT = "default"

var sortableProperties = []string{"maintainer", "version", "name", DEFAULT}
var sortTypes = []string{"asc", "desc", DEFAULT}
var filterableProperties = []string{"name", "maintainer", DEFAULT}

const (
	pageKey     = "page"
	sizeKey     = "size"
	sortByKey   = "sortBy"
	sortKey     = "sort"
	filterByKey = "filterBy"
	filterKey   = "filter"
)

func GetPackageListHandle(w http.ResponseWriter, req *http.Request) {

	packages := PackageList()
	params, err := query.
		New(req).
		OptionalInt(pageKey, 0).
		OptionalInt(sizeKey, 50).
		OptionalEnum(sortByKey, sortableProperties, DEFAULT).
		OptionalEnum(sortKey, sortTypes, DEFAULT).
		OptionalEnum(filterByKey, filterableProperties, DEFAULT).
		OptionalStr(filterKey, DEFAULT).
		MustComeTogheter([]string{sortKey, sortByKey}).
		MustComeTogheter([]string{filterKey, filterByKey}).
		Parse()

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf("{ error: \"%v\" }", err), 400)
		return
	}

	page := params.Ints[pageKey]
	pageSize := params.Ints[sizeKey]
	sort := params.Strings[sortKey]
	sortBy := params.Strings[sortByKey]
	filter := params.Strings[filterKey]
	filterBy := params.Strings[filterByKey]

	etag := fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v", LastModified().UTC().String(), page, pageSize, sort, sortBy, filter, filterBy)
	if serverlib.ApplyCacheHeaders(etag, &w, req) {
		// Response was cached and already sent
		return
	}

	packages = filterPackages(packages, filter, filterBy)
	packages = sortPackages(packages, sort, sortBy)
	packages = computePage(packages, page, pageSize)

	json, err := json.Marshal(packages)
	if err != nil {
		log.Printf("Could not marshal to json. Setting response 500.\n%v\n", err)
		w.WriteHeader(500)
	}

	serverlib.SendJson(&w, json)
}

func computePage(packages []types.PackageInfo, page, pageSize int) []types.PackageInfo {
	startIndex := page * pageSize
	endIndex := startIndex + pageSize

	if len(packages) < startIndex {
		return make([]types.PackageInfo, 0)
	}

	if len(packages) < endIndex {
		endIndex = len(packages)
	}

	return packages[startIndex:endIndex]
}

func filterPackages(packages []types.PackageInfo, filter, filterBy string) []types.PackageInfo {
	filterByFunc := func(matches func(*types.PackageInfo) bool) []types.PackageInfo {
		out := make([]types.PackageInfo, 0)
		for _, pkg := range packages {
			if matches(&pkg) {
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

type PackageListWrapper []types.PackageInfo

func (w PackageListWrapper) Len() int {
	return len(w)
}

func (w PackageListWrapper) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

type SortByName struct{ PackageListWrapper }
type SortByMaintainer struct{ PackageListWrapper }
type SortByVersion struct{ PackageListWrapper }

type SortByNameDesc struct{ PackageListWrapper }
type SortByMaintainerDesc struct{ PackageListWrapper }
type SortByVersionDesc struct{ PackageListWrapper }

func (s SortByName) Less(i, j int) bool {
	return strings.Compare(s.PackageListWrapper[i].Name, s.PackageListWrapper[j].Name) < 0
}

func (s SortByMaintainer) Less(i, j int) bool {
	return strings.Compare(s.PackageListWrapper[i].Maintainer, s.PackageListWrapper[j].Maintainer) < 0
}

func (s SortByVersion) Less(i, j int) bool {
	return strings.Compare(s.PackageListWrapper[i].Version, s.PackageListWrapper[j].Version) < 0
}

func (s SortByNameDesc) Less(i, j int) bool {
	return strings.Compare(s.PackageListWrapper[i].Name, s.PackageListWrapper[j].Name) > 0
}

func (s SortByMaintainerDesc) Less(i, j int) bool {
	return strings.Compare(s.PackageListWrapper[i].Maintainer, s.PackageListWrapper[j].Maintainer) > 0
}

func (s SortByVersionDesc) Less(i, j int) bool {
	return strings.Compare(s.PackageListWrapper[i].Version, s.PackageListWrapper[j].Version) > 0
}

func sortPackages(packages []types.PackageInfo, sortType, sortBy string) []types.PackageInfo {
	if strings.Compare(sortType, DEFAULT) == 0 {
		return packages
	}

	clone := make([]types.PackageInfo, 0)
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
