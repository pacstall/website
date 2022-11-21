package psapi

import (
	"fmt"
	"math"
	"net/http"

	"pacstall.dev/webserver/server"
	"pacstall.dev/webserver/types/pac/parser"
	"pacstall.dev/webserver/server/query"
	"pacstall.dev/webserver/types/pac/pacstore"
	"pacstall.dev/webserver/types/pac"
)

type packageListPage struct {
	Page     int           `json:"page"`
	Size     int           `json:"size"`
	Sort     string        `json:"sort"`
	SortBy   string        `json:"sortBy"`
	Filter   string        `json:"filter"`
	FilterBy string        `json:"filterBy"`
	Total    int           `json:"total"`
	LastPage int           `json:"lastPage"`
	Data     []*pac.Script `json:"data"`
}

func GetPacscriptListHandle(w http.ResponseWriter, req *http.Request) {

	packages := pacstore.GetAll().ToSlice()
	params, err := query.
		New(req).
		OptionalInt(parser.PageKey, 0).
		OptionalInt(parser.SizeKey, 50).
		OptionalEnum(parser.SortByKey, parser.SortableProperties, parser.DEFAULT).
		OptionalEnum(parser.SortKey, parser.SortTypes, parser.DEFAULT).
		OptionalEnum(parser.FilterByKey, parser.FilterableProperties, parser.DEFAULT).
		OptionalStr(parser.FilterKey, parser.DEFAULT).
		MustComeTogheter([]string{parser.SortKey, parser.SortByKey}).
		MustComeTogheter([]string{parser.FilterKey, parser.FilterByKey}).
		Parse()

	if err != nil {
		// w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	page := params.Ints[parser.PageKey]
	pageSize := params.Ints[parser.SizeKey]
	sort := params.Strings[parser.SortKey]
	sortBy := params.Strings[parser.SortByKey]
	filter := params.Strings[parser.FilterKey]
	filterBy := params.Strings[parser.FilterByKey]

	etag := fmt.Sprintf("%v-%v-%v-%v-%v-%v-%v", pacstore.LastModified().UTC().String(), page, pageSize, sort, sortBy, filter, filterBy)
	if server.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	packages = parser.FilterPackages(packages, filter, filterBy)
	packages = parser.SortPackages(packages, sort, sortBy)
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

	server.Json(w, result)
}

func computePage(packages []*pac.Script, page, pageSize int) []*pac.Script {
	startIndex := page * pageSize
	endIndex := startIndex + pageSize

	if len(packages) < startIndex {
		return make([]*pac.Script, 0)
	}

	if len(packages) < endIndex {
		endIndex = len(packages)
	}

	return packages[startIndex:endIndex]
}
