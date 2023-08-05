package repology

import (
	"errors"
	"fmt"
	"strings"

	"pacstall.dev/webserver/model"
	"pacstall.dev/webserver/types/list"
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
	_, projectName := parseRepologyFilter(search[0])

	query := db.Where("project_name = ?", projectName).Where(fmt.Sprintf("%v = ?", model.RepologyProjectProviderColumns.Active), true)
	for _, filter := range search[1:] {
		filterName, filterValue := parseRepologyFilter(filter)
		column, ok := repologyFilterToColumn[filterName]
		if !ok {
			return result, fmt.Errorf("invalid filter '%v'", filterName)
		}

		query = query.Where(fmt.Sprintf("%v = ?", column), filterValue)
	}

	var results []model.RepologyProjectProvider
	if err := query.Order("version desc").Find(&results).Error; err != nil || len(results) == 0 {
		return result, errors.Join(errors.New("failed to fetch repology project"), err)
	}

	results = sortByStatus(results)
	result = results[0]

	return result, nil
}

var repologyStatusPriority = map[string]int{
	"newest":    0,
	"rolling":   1,
	"devel":     3,
	"legacy":    4,
	"outdated":  5,
	"unique":    6,
	"noscheme":  7,
	"incorrect": 7,
	"untrusted": 7,
	"ignored":   7,
}

func sortByStatus(projects []model.RepologyProjectProvider) []model.RepologyProjectProvider {
	return list.From(projects).SortBy(func(p1, p2 model.RepologyProjectProvider) bool {
		return repologyStatusPriority[p1.Status] < repologyStatusPriority[p2.Status]
	}).ToSlice()
}
