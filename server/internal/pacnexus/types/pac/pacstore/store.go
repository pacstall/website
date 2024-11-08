package pacstore

import (
	"time"

	"pacstall.dev/webserver/internal/pacnexus/types/pac"
	"pacstall.dev/webserver/pkg/common/array"
)

var lastModified time.Time
var loadedPacscripts []*pac.Script

func FindByName(name string) (*pac.Script, error) {
	return array.FindBy(loadedPacscripts, func(p *pac.Script) bool {
		return p.PackageName == name
	})
}

func FindByMaintainer(maintainer string) (*pac.Script, error) {
	return array.FindBy(loadedPacscripts, func(p *pac.Script) bool {
		_, err := array.FindBy(p.Maintainers, func(s string) bool {
			return s == maintainer
		})

		return err != nil
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
