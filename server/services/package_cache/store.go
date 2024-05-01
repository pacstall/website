package pkgcache

import (
	"time"

	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
)

type PackageCacheService struct {
	lastModified     time.Time
	loadedPacscripts []*pac.Script
}

func New() *PackageCacheService {
	return &PackageCacheService{}
}

func (s *PackageCacheService) FindByName(name string) (*pac.Script, error) {
	return array.FindBy(s.loadedPacscripts, func(p *pac.Script) bool {
		return p.PackageName == name
	})
}

func (s *PackageCacheService) FindByMaintainer(maintainer string) (*pac.Script, error) {
	return array.FindBy(s.loadedPacscripts, func(p *pac.Script) bool {
		_, err := array.FindBy(p.Maintainers, func(s string) bool {
			return s == maintainer
		})

		return err != nil
	})
}

func (s *PackageCacheService) GetAll() []*pac.Script {
	return array.Clone(s.loadedPacscripts)
}

func (s *PackageCacheService) LastModified() time.Time {
	return s.lastModified
}

func (s *PackageCacheService) Update(scripts []*pac.Script) {
	s.lastModified = time.Now()
	s.loadedPacscripts = scripts
}
