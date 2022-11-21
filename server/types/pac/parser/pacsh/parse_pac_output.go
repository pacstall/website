package pacsh

import (
	"fmt"
	"strings"

	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

var ParsePacOutput = parseOutput
var PacstallCVars []string = []string{"name", "pkgname", "maintainer", "description", "url", "gives", "hash", "version"}
var PacstallCArrays []string = []string{"depends", "breaks", "replace", "build_depends"}
var PacstallCMaps []string = []string{"optdepends", "pacdeps", "patch", "ppa", "repology"}

const (
	nameIdx = iota
	pkgnameIdx
	maintainerIdx
	descriptionIdx
	urlIdx
	givesIdx
	hashIdx
	versionIdx
	dependsIdx
	breaksIdx
	replaceIdx
	buildDependsIdx
	optdependsIdx
	pacdepsIdx
	patchIdx
	ppaIdx
	repologyIdx
)

func parseSubcategory(category string) []string {
	return list.From(strings.Split(category, "+  +++")).Map(func(s string) string {
		return strings.TrimSpace(s)
	}).Filter(list.Not(""))
}

func parseOutput(data []byte) (out pac.Script, err error) {
	content := string(data)

	categories := list.From(strings.Split(content, "++++")).Map(func(s string) string { return strings.TrimSpace(s) }).ToSlice()[1:]
	name := categories[nameIdx]
	packageName := categories[pkgnameIdx]
	maintainer := categories[maintainerIdx]
	description := categories[descriptionIdx]
	url := categories[urlIdx]
	gives := categories[givesIdx]
	hash := categories[hashIdx]
	version := categories[versionIdx]
	runtimeDependencies := categories[dependsIdx]
	breaks := categories[breaksIdx]
	replace := categories[replaceIdx]
	buildDependencies := categories[buildDependsIdx]
	optionalDependencies := categories[optdependsIdx]
	pacstallDependencies := categories[pacdepsIdx]
	patch := categories[patchIdx]
	ppa := categories[ppaIdx]
	repology := categories[repologyIdx]

	hashPtr := &hash
	if hash == "" {
		hashPtr = nil
	}

	if len(strings.TrimSpace(version)) == 0 {
		return out, fmt.Errorf("version is empty")
	}

	out = pac.Script{
		Name:                 name,
		PackageName:          packageName,
		Maintainer:           maintainer,
		Description:          description,
		URL:                  url,
		Gives:                gives,
		Hash:                 hashPtr,
		Version:              version,
		RuntimeDependencies:  strings.Fields(runtimeDependencies),
		Breaks:               strings.Fields(breaks),
		Replace:              strings.Fields(replace),
		BuildDependencies:    strings.Fields(buildDependencies),
		OptionalDependencies: parseSubcategory(optionalDependencies),
		PacstallDependencies: parseSubcategory(pacstallDependencies),
		PPA:                  parseSubcategory(ppa),
		Patch:                parseSubcategory(patch),
		RequiredBy:           make([]string, 0),
		Repology:             parseSubcategory(repology),
		LatestVersion:        nil,
		UpdateStatus:         pac.UpdateStatus.Unknown,
	}

	if out.PackageName == "" {
		out.PackageName = out.Name
	}

	out.PrettyName = getPrettyName(out)

	return
}
