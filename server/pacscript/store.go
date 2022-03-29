package pacscript

import (
	"time"

	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/list"
)

type PackageList struct {
	list.List[*types.Pacscript]
}

var lastModified time.Time
var loadedPackages list.List[*types.Pacscript]

func (l PackageList) FindByName(name string) (*types.Pacscript, error) {
	return l.FindBy(func(pi *types.Pacscript) bool {
		return pi.Name == name
	})
}

func (l PackageList) FindByMaintainer(maintainer string) (*types.Pacscript, error) {
	return l.FindBy(func(pi *types.Pacscript) bool {
		return pi.Maintainer == maintainer
	})
}
