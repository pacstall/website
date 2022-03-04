package pacscript

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"pacstall.dev/website/listener"
	"pacstall.dev/website/listener/query"
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

func GetPackageHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name, ok := params["name"]
	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	if listener.ApplyHeaders(fmt.Sprintf("%v-%v", LastModified().UTC().String(), name), w, req) {
		return // req is cached
	}

	for _, pkg := range PackageList() {
		if strings.Compare(pkg.Name, name) == 0 {
			listener.Json(w, pkg)
			return
		}
	}

	w.WriteHeader(404)
}

type packageListPage struct {
	Page     int                  `json:"page"`
	Size     int                  `json:"size"`
	Sort     string               `json:"sort"`
	SortBy   string               `json:"sortBy"`
	Filter   string               `json:"filter"`
	FilterBy string               `json:"filterBy"`
	Total    int                  `json:"total"`
	LastPage int                  `json:"lastPage"`
	Data     []*types.PackageInfo `json:"data"`
}

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
	if listener.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	packages = filterPackages(packages, filter, filterBy)
	packages = sortPackages(packages, sort, sortBy)
	found := len(packages)
	packages = computePage(packages, page, pageSize)

	result := packageListPage{
		Page:     page,
		Size:     pageSize,
		Sort:     sort,
		SortBy:   sortBy,
		Filter:   filter,
		FilterBy: filterBy,
		Total:    found,
		LastPage: int(math.Floor(float64(found) / float64(pageSize))),
		Data:     packages,
	}

	listener.Json(w, result)
}

func GetPackagesRequiredByHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name, ok := params["name"]

	etag := fmt.Sprintf("%v-%v", LastModified().UTC().String(), name)
	if listener.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	allPackages := PackageList()

	var pacpkg *types.PackageInfo
	for _, it := range allPackages {
		if name == it.Name {
			pacpkg = it
			break
		}
	}

	if pacpkg == nil {
		w.WriteHeader(404)
		return
	}

	requiredBy := make([]*types.PackageInfo, 0)
	for _, pkg := range allPackages {
		for _, requiredByName := range pacpkg.RequiredBy {
			if strings.Compare(pkg.Name, requiredByName) == 0 {
				requiredBy = append(requiredBy, pkg)
				break
			}
		}
	}

	listener.Json(w, requiredBy)
}

type pacPackageOrExternal interface{}

type packageDependencies struct {
	RuntimeDependencies  []pacPackageOrExternal `json:"runtimeDependencies"`
	BuildDependencies    []pacPackageOrExternal `json:"buildDependencies"`
	OptionalDependencies []pacPackageOrExternal `json:"optionalDependencies"`
	PacstallDependencies []pacPackageOrExternal `json:"pacstallDependencies"`
}

func GetPackageDependenciesHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name, ok := params["name"]

	etag := fmt.Sprintf("%v-%v", LastModified().UTC().String(), name)
	if listener.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	allPackages := PackageList()

	var pacpkg *types.PackageInfo
	for _, it := range allPackages {
		if name == it.Name {
			pacpkg = it
			break
		}
	}

	if pacpkg == nil {
		w.WriteHeader(404)
		return
	}

	buildDependencies := make([]pacPackageOrExternal, 0)
	for _, pkg := range pacpkg.BuildDependencies {
		if found := findPackageInList(pkg, allPackages); found != nil {
			buildDependencies = append(buildDependencies, found)
		} else {
			buildDependencies = append(buildDependencies, pkg)
		}
	}

	runtimeDependencies := make([]pacPackageOrExternal, 0)
	for _, pkg := range pacpkg.RuntimeDependencies {
		if found := findPackageInList(pkg, allPackages); found != nil {
			runtimeDependencies = append(runtimeDependencies, found)
		} else {
			runtimeDependencies = append(runtimeDependencies, pkg)
		}
	}

	optionalDependencies := make([]pacPackageOrExternal, 0)
	for _, pkg := range pacpkg.OptionalDependencies {
		if found := findPackageInList(pkg, allPackages); found != nil {
			optionalDependencies = append(optionalDependencies, found)
		} else {
			optionalDependencies = append(optionalDependencies, pkg)
		}
	}

	pacstallDependencies := make([]pacPackageOrExternal, 0)
	for _, pkg := range pacpkg.PacstallDependencies {
		if found := findPackageInList(pkg, allPackages); found != nil {
			pacstallDependencies = append(pacstallDependencies, found)
		} else {
			pacstallDependencies = append(pacstallDependencies, pkg)
		}
	}

	response := packageDependencies{
		RuntimeDependencies:  runtimeDependencies,
		BuildDependencies:    buildDependencies,
		OptionalDependencies: optionalDependencies,
		PacstallDependencies: pacstallDependencies,
	}

	listener.Json(w, response)
}

func findPackageInList(name string, packages []*types.PackageInfo) *types.PackageInfo {
	for _, pkg := range packages {
		if strings.Compare(pkg.Name, name) == 0 {
			return pkg
		}
	}

	return nil
}

func computePage(packages []*types.PackageInfo, page, pageSize int) []*types.PackageInfo {
	startIndex := page * pageSize
	endIndex := startIndex + pageSize

	if len(packages) < startIndex {
		return make([]*types.PackageInfo, 0)
	}

	if len(packages) < endIndex {
		endIndex = len(packages)
	}

	return packages[startIndex:endIndex]
}

func filterPackages(packages []*types.PackageInfo, filter, filterBy string) []*types.PackageInfo {
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

type PackageListWrapper []*types.PackageInfo

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

func sortPackages(packages []*types.PackageInfo, sortType, sortBy string) []*types.PackageInfo {
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
