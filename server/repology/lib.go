package repology

import (
	"pacstall.dev/webserver/types/pac"
)

func Sync(script *pac.Script) error {
	if len(script.Repology) == 0 {
		return nil
	}

	project, err := findRepologyProject(script.Repology)
	if err != nil {
		return err
	}

	return updateScriptVersion(project, script)
}
