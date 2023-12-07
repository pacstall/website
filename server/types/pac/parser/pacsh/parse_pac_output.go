package pacsh

import (
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
)

var ParsePacOutput = parseOutput
var PacstallCVars []string = []string{"name", "pkgname", "maintainer", "pkgdesc", "url", "gives", "hash", "pkgver"}
var PacstallCArrays []string = []string{}
var PacstallCMaps []string = []string{"depends", "breaks", "replace", "makedepends", "optdepends", "pacdeps", "patch", "ppa", "repology"}

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
	subcategories := strings.Split(category, "+  +++")
	for i, subcategory := range subcategories {
		subcategories[i] = strings.TrimSpace(subcategory)
	}

	return array.Filter(subcategories, func(it *array.Iterator[string]) bool {
		return len(it.Value) > 0
	})
}

func parseOutput(data []byte) (out pac.Script, err error) {
	content := string(data)

	categories := array.Map(strings.Split(content, "++++"), func(it *array.Iterator[string]) string { return strings.TrimSpace(it.Value) })[1:]
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
		if strings.HasSuffix(name, "-git") {
			version = "git"
		} else {
			return out, errorx.IllegalArgument.New("expected version to be non-empty but got: %v", version)
		}
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
		RuntimeDependencies:  parseSubcategory(runtimeDependencies),
		Breaks:               parseSubcategory(breaks),
		Replace:              parseSubcategory(replace),
		BuildDependencies:    parseSubcategory(buildDependencies),
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
