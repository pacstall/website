package repology

import (
	"errors"
	"fmt"
	"strings"

	"pacstall.dev/webserver/model"
)

func parseRepologyFilter(filter string) (string, string) {
	idx := strings.Index(filter, ":")
	return strings.TrimSpace(filter[:idx]), strings.TrimSpace(filter[idx+1:])
}

var repologyFilterToColumn = map[string]string{
	"repo":        model.RepologyProjectProviderColumns.Repository,
	"subrepo":     model.RepologyProjectProviderColumns.SubRepository,
	"status":      model.RepologyProjectProviderColumns.Status,
	"srcname":     model.RepologyProjectProviderColumns.SourceName,
	"binname":     model.RepologyProjectProviderColumns.BinaryName,
	"version":     model.RepologyProjectProviderColumns.Version,
	"origversion": model.RepologyProjectProviderColumns.OriginalVersion,
	"visiblename": model.RepologyProjectProviderColumns.VisibleName,
	"summary":     model.RepologyProjectProviderColumns.Summary,
}

func findRepologyProject(search []string) (model.RepologyProjectProvider, error) {
	var result model.RepologyProjectProvider

	if len(search) == 0 {
		return result, fmt.Errorf("no search terms provided")
	}

	db := model.Instance()

	query := db.Where("project_name = ?", search[0])
	for _, filter := range search[1:] {
		filterName, filterValue := parseRepologyFilter(filter)
		column, ok := repologyFilterToColumn[filterName]
		if !ok {
			return result, fmt.Errorf("invalid filter '%v'", filterName)
		}

		query = query.Where(fmt.Sprintf("%v = ?", column), filterValue)
	}

	if err := query.Order("version desc").First(&result).Error; err != nil {
		return result, errors.Join(errors.New("failed to fetch repology project"), err)
	}

	return result, nil
}
