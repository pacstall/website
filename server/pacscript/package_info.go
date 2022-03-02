package pacscript

import (
	"encoding/json"
	"fmt"
	"strings"

	"pacstall.dev/website/types"
	"pacstall.dev/website/types/list"
)

type rawPackageInfo struct {
	Name                 string   `yaml:"name"`
	Version              string   `yaml:"version"`
	PackageName          string   `yaml:"packageName"`
	Maintainer           string   `yaml:"maintainer"`
	Description          string   `yaml:"description"`
	URL                  string   `yaml:"url"`
	RuntimeDependencies  string   `yaml:"runtimeDependencies"`
	BuildDependencies    string   `yaml:"buildDependencies"`
	OptionalDependencies []string `yaml:"optionalDependencies"`
	Breaks               string   `yaml:"breaks"`
	Gives                string   `yaml:"gives"`
	Replace              string   `yaml:"replace"`
	Hash                 string   `yaml:"hash"`
	PPA                  []string `yaml:"ppa"`
	PacstallDependencies []string `yaml:"pacstallDependencies"`
	Patch                []string `yaml:"patch"`
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
	out["version"] = "version"
	return out
}

func makePacArrays() map[jsonName]bashName {
	out := make(map[jsonName]bashName)
	out["runtimeDependencies"] = "depends"
	out["breaks"] = "breaks"
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

func removeDebianCheckForAnAnimeGame(script string, animeGameChecked *bool) string {
	if *animeGameChecked {
		return script
	}

	if !strings.Contains(script, "name=an-anime-game-launcher-bin") {
		return script
	}

	*animeGameChecked = true

	if strings.Index(script, "if") != 0 {
		return script
	}

	debianCheckEnd := strings.Index(script, "fi")
	if debianCheckEnd == -1 {
		return script
	}

	return script[debianCheckEnd+len("fi"):]
}

func buildYamlScript(header string) string {
	// TODO: remove after `preinstall` gets implemented
	animeGameChecked := false
	script := removeDebianCheckForAnAnimeGame(header, &animeGameChecked) + "\n"
	script = script + "echo ''\n"

	for jsonName, bashName := range pacstallVars {
		script += fmt.Sprintf("echo \"%v: >\"", jsonName) + "\n"
		script += fmt.Sprintf("echo \"  $%v\"", bashName) + "\n"
	}

	for jsonName, bashName := range pacstallArrays {
		script += fmt.Sprintf("echo \"%v: >\"", jsonName) + "\n"
		script += fmt.Sprintf("echo \"  $%v\"", bashName) + "\n"
	}

	mapsPartialScript := make([]string, 0)
	for jsonName, bashName := range pacstallMaps {
		partial := ""

		partial += fmt.Sprintf(`echo %v:`, jsonName) + "\n"
		partial += fmt.Sprintf("for val in ${%v[@]}\n", bashName)
		partial += "do\n"
		partial += "echo \"  - >\"" + "\n"
		partial += "echo \"    $val\"" + "\n"
		partial += "done\n"

		mapsPartialScript = append(mapsPartialScript, partial)
	}

	script += strings.Join(mapsPartialScript, "\n")

	return script
}

func RepairPackageInfo(pkg *types.PackageInfo) error {
	bytes, err := json.Marshal(pkg)
	if err != nil {
		return err
	}

	content := string(bytes)

	content = strings.ReplaceAll(content, `\n`, "")
	content = strings.ReplaceAll(content, `[""]`, "[]")
	content = strings.ReplaceAll(content, `":null`, `":[]`)

	if err = json.Unmarshal([]byte(content), &pkg); err != nil {
		return err
	}

	return nil
}

func (rp rawPackageInfo) toPackageInfo() types.PackageInfo {
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
		RequiredBy:           make([]string, 0),
		Version:              rp.Version,
	}

	return out
}

func computeRequiredBy(pkgs []types.PackageInfo) []types.PackageInfo {
	strEq := func(s1, s2 string) bool { return s1 == s2 }
	pickBeforeColon := func(arr []string) []string {
		out := make([]string, len(arr))
		for _, it := range arr {
			out = append(out, strings.Split(it, ":")[0])
		}
		return out
	}

	for idx, pkg := range pkgs {
		for otherIdx, otherPkg := range pkgs {
			if idx == otherIdx {
				continue
			}

			allDeps := make([]string, 0)
			allDeps = append(allDeps, pickBeforeColon(otherPkg.BuildDependencies)...)
			allDeps = append(allDeps, pickBeforeColon(otherPkg.RuntimeDependencies)...)
			allDeps = append(allDeps, pickBeforeColon(otherPkg.OptionalDependencies)...)
			allDeps = append(allDeps, pickBeforeColon(otherPkg.PacstallDependencies)...)
			allDepsList := list.StrList(allDeps)

			if allDepsList.Contains(pkg.Name, strEq) ||
				allDepsList.Contains(pkg.Gives, strEq) ||
				allDepsList.Contains(pkg.PackageName, strEq) {
				pkg.RequiredBy = append(pkg.RequiredBy, otherPkg.Name)
			}
		}
	}

	return pkgs
}
