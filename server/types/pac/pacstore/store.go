package pacstore

import (
	"time"

	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
)

var lastModified time.Time
var loadedPacscripts []*pac.Script

func FindByName(name string) (*pac.Script, error) {
	return array.FindBy(loadedPacscripts, func(p *pac.Script) bool {
		return p.Name == name
	})
}

func FindByMaintainer(maintainer string) (*pac.Script, error) {
	return array.FindBy(loadedPacscripts, func(p *pac.Script) bool {
		return p.Maintainer == maintainer
	})
}

func GetAll() []*pac.Script {
	return array.Clone(loadedPacscripts)
}

func LastModified() time.Time {
	return lastModified
}

func Update(scripts []*pac.Script) {
	lastModified = time.Now()
	loadedPacscripts = scripts
}
