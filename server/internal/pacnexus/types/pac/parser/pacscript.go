package parser

import (
	"pacstall.dev/webserver/internal/pacnexus/types/pac"
	"pacstall.dev/webserver/pkg/common/array"
)

func computeRequiredBy(script *pac.Script, scripts []*pac.Script) {
	pickName := func(it *array.Iterator[pac.ArchDistroString]) string {
		return it.Value.Value
	}

	script.RequiredBy = make([]string, 0)
	for _, otherScript := range scripts {
		otherScriptDependencies := array.SwitchMap(otherScript.PacstallDependencies, pickName)
		if array.Contains(otherScriptDependencies, array.Is(script.PackageName)) {
			script.RequiredBy = append(script.RequiredBy, otherScript.PackageName)
		}
	}
}
