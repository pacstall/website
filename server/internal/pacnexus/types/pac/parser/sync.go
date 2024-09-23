package parser

import (
	"strings"

	"github.com/hashicorp/go-version"
	"pacstall.dev/webserver/internal/pacnexus/types/pac"
	"pacstall.dev/webserver/pkg/common/types"
)

func compareNonStandardVersion(current, latest string) pac.UpdateStatusValue {
	result := strings.Compare(current, latest)

	const CMP_EQUAL = 0
	const CMP_GREATER = 1

	if result == CMP_EQUAL || result == CMP_GREATER {
		return pac.UpdateStatus.Latest
	}

	return pac.UpdateStatus.Major
}

func UpdateScriptVersion(project types.RepologyApiProject, script *pac.Script) (err error) {
	script.LatestVersion = &project.Version

	if *script.LatestVersion == script.Version {
		script.UpdateStatus = pac.UpdateStatus.Latest
		return
	}

	current, err := version.NewVersion(script.Version)
	if err != nil {
		err = nil
		script.UpdateStatus = compareNonStandardVersion(script.Version, *script.LatestVersion)
		return
	}

	latest, err := version.NewVersion(*script.LatestVersion)
	if err != nil {
		err = nil
		script.UpdateStatus = compareNonStandardVersion(script.Version, *script.LatestVersion)
		return
	}

	if latest.LessThanOrEqual(current) {
		script.UpdateStatus = pac.UpdateStatus.Latest
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

func GetUpdateStatus(current, latest string) pac.UpdateStatusValue {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	currentParts := strings.Split(current, ".")
	latestParts := strings.Split(latest, ".")
	minParts := min(len(currentParts), len(latestParts))
	versionDiff := 0
	for i := 0; i < minParts; i++ {
		if currentParts[i] < latestParts[i] {
			break
		} else if currentParts[i] > latestParts[i] {
			return pac.UpdateStatus.Latest
		}

		versionDiff += 1
	}

	if minParts == 0 {
		return pac.UpdateStatus.Unknown
	}

	if versionDiff == minParts {
		return pac.UpdateStatus.Latest
	}

	if versionDiff == 1 {
		return pac.UpdateStatus.Major
	}

	if versionDiff == 2 {
		return pac.UpdateStatus.Minor
	}

	return pac.UpdateStatus.Patch
}
