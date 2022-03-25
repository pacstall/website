package pshttphandle

import (
	"fmt"
	"math"
	"net/http"

	"pacstall.dev/website/listener"
	"pacstall.dev/website/listener/query"
	"pacstall.dev/website/pacscript"
	"pacstall.dev/website/types"
)

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

	packages := pacscript.PackageList()
	params, err := query.
		New(req).
		OptionalInt(pacscript.PageKey, 0).
		OptionalInt(pacscript.SizeKey, 50).
		OptionalEnum(pacscript.SortByKey, pacscript.SortableProperties, pacscript.DEFAULT).
		OptionalEnum(pacscript.SortKey, pacscript.SortTypes, pacscript.DEFAULT).
		OptionalEnum(pacscript.FilterByKey, pacscript.FilterableProperties, pacscript.DEFAULT).
		OptionalStr(pacscript.FilterKey, pacscript.DEFAULT).
		MustComeTogheter([]string{pacscript.SortKey, pacscript.SortByKey}).
		MustComeTogheter([]string{pacscript.FilterKey, pacscript.FilterByKey}).
		Parse()

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf("{ error: \"%v\" }", err), 400)
		return
	}

	page := params.Ints[pacscript.PageKey]
	pageSize := params.Ints[pacscript.SizeKey]
	sort := params.Strings[pacscript.SortKey]
	sortBy := params.Strings[pacscript.SortByKey]
	filter := params.Strings[pacscript.FilterKey]
	filterBy := params.Strings[pacscript.FilterByKey]

	etag := fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v", pacscript.LastModified().UTC().String(), page, pageSize, sort, sortBy, filter, filterBy)
	if listener.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	packages = pacscript.FilterPackages(packages, filter, filterBy)
	packages = pacscript.SortPackages(packages, sort, sortBy)
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
