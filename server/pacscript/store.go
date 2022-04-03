package pacscript

import (
	"time"

	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

type PacscriptList struct {
	list.List[*pac.Script]
}

var lastModified time.Time
var loadedPacscripts list.List[*pac.Script]

func (l PacscriptList) FindByName(name string) (*pac.Script, error) {
	return l.FindBy(func(pi *pac.Script) bool {
		return pi.Name == name
	})
}

func (l PacscriptList) FindByMaintainer(maintainer string) (*pac.Script, error) {
	return l.FindBy(func(pi *pac.Script) bool {
		return pi.Maintainer == maintainer
	})
}
