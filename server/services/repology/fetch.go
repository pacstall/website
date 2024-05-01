package repology

import (
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/model"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/repository"
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

func (s *RepologyService) findRepologyProject(search []string) (repository.RepologyProjectProvider, error) {
	var result repository.RepologyProjectProvider

	if len(search) == 0 {
		return result, errorx.IllegalArgument.New("no search terms provided")
	}

	_, projectName := parseRepologyFilter(search[0])

	filters := make(map[string]interface{})
	for _, filter := range search[1:] {
		filterName, filterValue := parseRepologyFilter(filter)
		column, ok := repologyFilterToColumn[filterName]
		if !ok {
			return result, errorx.IllegalArgument.New("invalid filter '%v'", filterName)
		}

		filters[column] = filterValue
	}

	results, err := s.repologyProjectProviderRepository.FindAllWhereProjectNameAndFiltersSortedByVersionDesc(projectName, filters)
	if err != nil || len(results) == 0 {
		return result, errorx.Decorate(err, "failed to fetch repology project")
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

func sortByStatus(projects []repository.RepologyProjectProvider) []repository.RepologyProjectProvider {
	return array.SortBy(array.Clone(projects), func(p1, p2 repository.RepologyProjectProvider) bool {
		return repologyStatusPriority[p1.GetStatus()] < repologyStatusPriority[p2.GetStatus()]
	})
}
