package pacscript

import (
	"strings"

	"pacstall.dev/website/types"
)

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
