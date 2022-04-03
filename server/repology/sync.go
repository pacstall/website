package repology

import (
	"github.com/hashicorp/go-version"
	"pacstall.dev/webserver/types/pac"
)

func syncToPacscript(project repologyProject, script *pac.Script) (err error) {
	script.PrettyName = project.PrettyName
	script.LatestVersion = project.Version

	if script.LatestVersion == script.Version {
		script.UpdateStatus = pac.UpdateStatus.Latest
		return
	}

	current, err := version.NewVersion(script.Version)
	if err != nil {
		err = nil
		script.UpdateStatus = pac.UpdateStatus.Unknown
		return
	}

	latest, err := version.NewVersion(script.LatestVersion)
	if err != nil {
		err = nil
		script.UpdateStatus = pac.UpdateStatus.Unknown
		return
	}

	currentVersionParts := current.Segments64()
	latestVersionParts := latest.Segments64()

	script.UpdateStatus = pac.UpdateStatus.Minor
	if currentVersionParts[0] < latestVersionParts[0] {
		script.UpdateStatus = pac.UpdateStatus.Major
		return
	}

	if len(currentVersionParts) < 2 && len(latestVersionParts) < 2 {
		return
	}

	script.UpdateStatus = pac.UpdateStatus.Patch
	if currentVersionParts[1] < latestVersionParts[1] {
		script.UpdateStatus = pac.UpdateStatus.Minor
		return
	}

	if len(currentVersionParts) < 3 && len(latestVersionParts) < 3 {
		return
	}

	script.UpdateStatus = pac.UpdateStatus.Latest
	if currentVersionParts[2] < latestVersionParts[2] {
		script.UpdateStatus = pac.UpdateStatus.Patch
		return
	}

	return
}
