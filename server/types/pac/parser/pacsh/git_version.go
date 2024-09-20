package pacsh

import (
	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/parser/pacsh/internal"
)

func ApplyGitVersion(p *pac.Script) error {
	sources := internal.NewGitSources(array.SwitchMap(p.Source, func(it *array.Iterator[pac.ArchDistroString]) string {
		return it.Value.Value
	}))

	version, err := sources.ParseGitPackageVersion()
	if err != nil {
		return errorx.Decorate(err, "failed to parse git version for package '%s'", p.PackageName)
	}

	if p.Epoch != "" {
		p.Version = p.Epoch + ":" + version + "-" + p.Release
	} else {
		p.Version = version + "-" + p.Release
	}

	return nil
}
