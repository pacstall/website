package pacstore

import (
	"time"

	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

type PacscriptList struct {
	list.List[*pac.Script]
}

var lastModified time.Time
var loadedPacscripts PacscriptList

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

func GetAll() PacscriptList {
	return PacscriptList{
		loadedPacscripts.Clone().ToSlice(),
	}
}

func LastModified() time.Time {
	return lastModified
}

func Update(scripts list.List[*pac.Script]) {
	lastModified = time.Now()
	loadedPacscripts = PacscriptList{
		scripts.Clone(),
	}
}
