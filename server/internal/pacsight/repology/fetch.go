package repology

import (
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/pkg/common/array"
	"pacstall.dev/webserver/pkg/common/types"
)

func parseRepologyFilter(filter string) (string, string) {
	idx := strings.Index(filter, ":")
	return strings.TrimSpace(filter[:idx]), strings.TrimSpace(filter[idx+1:])
}

const (
	RF_REPO        = "repo"
	RF_SUBREPO     = "subrepo"
	RF_STATUS      = "status"
	RF_SRCNAME     = "srcname"
	RF_BINNAME     = "binname"
	RF_VERSION     = "version"
	RF_ORIGVERSION = "origversion"
	RF_VISIBLENAME = "visiblename"
	RF_SUMMARY     = "summary"
)

func FindRepologyProject(projectsMap types.RepologyApiProjectSearchResponse, search []string) (types.RepologyApiProject, error) {
	if len(search) == 0 {
		return types.RepologyApiProject{}, errorx.IllegalArgument.New("no search terms provided")
	}

	_, projectName := parseRepologyFilter(search[0])

	projects, ok := projectsMap[projectName]
	if !ok {
		return types.RepologyApiProject{}, errorx.DataUnavailable.
			New("project not found").
			WithProperty(errorx.RegisterProperty("project"), projectName)
	}

	for _, filter := range search[1:] {
		field, value := parseRepologyFilter(filter)

		projects = array.Filter(projects, func(i *array.Iterator[types.RepologyApiProject]) bool {
			switch field {
			case RF_REPO:
				return i.Value.Repository == value
			case RF_SUBREPO:
				return i.Value.SubRepository != nil && *i.Value.SubRepository == value
			case RF_STATUS:
				return i.Value.Status == value
			case RF_SRCNAME:
				return i.Value.SourceName != nil && *i.Value.SourceName == value
			case RF_BINNAME:
				return i.Value.BinaryName != nil && *i.Value.BinaryName == value
			case RF_VERSION:
				return i.Value.Version == value
			case RF_ORIGVERSION:
				return i.Value.OriginalVersion == value
			case RF_VISIBLENAME:
				return i.Value.VisibleName != nil && *i.Value.VisibleName == value
			case RF_SUMMARY:
				return i.Value.Summary == value
			default:
				return false
			}
		})
	}

	projects = sortByStatus(projects)

	if len(projects) == 0 {
		return types.RepologyApiProject{}, errorx.IllegalArgument.New("no projects found")
	}

	return projects[0], nil
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

func sortByStatus(projects []types.RepologyApiProject) []types.RepologyApiProject {
	return array.SortBy(array.Clone(projects), func(p1, p2 types.RepologyApiProject) bool {
		return repologyStatusPriority[p1.Status] < repologyStatusPriority[p2.Status]
	})
}
