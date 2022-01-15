package pacpkgs

import (
	"fmt"
	"strings"

	"pacstall.dev/website/types"
)

type rawPackageInfo struct {
	Name                 string   `json:"name"`
	PackageName          string   `json:"packageName"`
	Maintainer           string   `json:"maintainer"`
	Description          string   `json:"description"`
	URL                  string   `json:"url"`
	RuntimeDependencies  string   `json:"runtimeDependencies"`
	BuildDependencies    string   `json:"buildDependencies"`
	OptionalDependencies []string `json:"optionalDependencies"`
	Breaks               string   `json:"breaks"`
	Gives                string   `json:"gives"`
	Replace              string   `json:"replace"`
	Hash                 string   `json:"hash"`
	PPA                  []string `json:"ppa"`
	PacstallDependencies []string `json:"pacstallDependencies"`
	Patch                []string `json:"patch"`
}

type jsonName = string
type bashName = string

var pacstallVars map[jsonName]bashName = makePacVars()
var pacstallArrays map[jsonName]bashName = makePacArrays()
var pacstallMaps map[jsonName]bashName = makePacMaps()

func makePacVars() map[jsonName]bashName {
	out := make(map[jsonName]bashName)
	out["name"] = "name"
	out["packageName"] = "pkgname"
	out["maintainer"] = "maintainer"
	out["description"] = "description"
	out["url"] = "url"
	out["gives"] = "gives"
	out["hash"] = "hash"
	return out
}

func makePacArrays() map[jsonName]bashName {
	out := make(map[jsonName]bashName)
	out["runtimeDependencies"] = "depends"
	out["breaks"] = "breaks"
	out["gives"] = "gives"
	out["replace"] = "replace"
	out["buildDependencies"] = "build_depends"
	return out
}

func makePacMaps() map[jsonName]bashName {
	out := make(map[jsonName]bashName)
	out["optionalDependencies"] = "optdepends"
	out["ppa"] = "ppa"
	out["pacstallDependencies"] = "pacdeps"
	out["patch"] = "patch"
	return out
}

func BuildJsonScript(header string) string {
	quote := func(what string) string { return fmt.Sprintf(`\"%v\"`, what) }

	script := header + "\n"

	script += "echo {\n"

	for jsonName, bashName := range pacstallVars {
		script += fmt.Sprintf(`echo"   %v: %v,"`, quote(jsonName), quote("$"+bashName)) + "\n"
	}

	for jsonName, bashName := range pacstallArrays {
		script += fmt.Sprintf(`echo "   %v: %v,"`, quote(jsonName), quote("$"+bashName)) + "\n"
	}

	mapsPartialScript := make([]string, 0)
	for jsonName, bashName := range pacstallMaps {
		partial := ""

		partial += fmt.Sprintf("echo -n \"   %v: [\"\n", quote(jsonName))
		partial += fmt.Sprintf("for val in ${%v[@]}\n", quote(bashName))
		partial += "do\n"
		partial += fmt.Sprintf("echo -n \"  %v,\"\n", quote("$val"))
		partial += "done\n"
		partial += "echo -n ]\n"

		mapsPartialScript = append(mapsPartialScript, partial)
	}

	script += strings.Join(mapsPartialScript, "echo \",\"\n")
	script += "echo \"\"\n"
	script += "echo }\n"

	script = strings.ReplaceAll(script, ", ]", " ]")

	return script
}

func (rp *rawPackageInfo) toPackageInfo() *types.PackageInfo {
	out := types.PackageInfo{
		Name:                 rp.Name,
		PackageName:          rp.PackageName,
		Maintainer:           rp.Maintainer,
		Description:          rp.Description,
		URL:                  rp.URL,
		RuntimeDependencies:  strings.Split(rp.RuntimeDependencies, " "),
		BuildDependencies:    strings.Split(rp.BuildDependencies, " "),
		Hash:                 rp.Hash,
		Breaks:               strings.Split(rp.Breaks, " "),
		Gives:                rp.Gives,
		Replace:              strings.Split(rp.Replace, " "),
		OptionalDependencies: rp.OptionalDependencies,
		PPA:                  rp.PPA,
		PacstallDependencies: rp.PacstallDependencies,
		Patch:                rp.Patch,
	}

	return &out
}
