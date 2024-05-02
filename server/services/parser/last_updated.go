package parser

import (
	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/service"
)

func (s *ParserService) setLastUpdatedAt(packages []*pac.Script, programsClonePath string) error {
	lastUpdatedTuples, err := s.packageLastUpdatedService.GetPackagesLastUpdated(programsClonePath)
	if err != nil {
		return errorx.Decorate(err, "failed to get package last updated tuples")
	}

	for _, pkg := range packages {
		if tuple, err := array.FindBy(lastUpdatedTuples, func(tuple service.PackageLastUpdatedTuple) bool {
			return tuple.PackageName == pkg.PackageName
		}); err == nil {
			pkg.LastUpdatedAt = tuple.LastUpdated
		} else {
			log.Warn("failed to set 'LastUpdatedAt' for package %#v. err: %+v", pkg, err)
		}
	}

	return nil
}
