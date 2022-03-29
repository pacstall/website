package pacscript

import (
	"encoding/json"
	"fmt"
	"strings"

	"pacstall.dev/webserver/types"
)

type rawPacscript struct {
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

func removeDebianCheck(script string) string {
	if !strings.Contains(script, "/etc/os-release)\" == \"debian\"") {
		return script
	}

	if strings.Index(script, "if") != 0 {
		return script
	}

	debianCheckEnd := strings.Index(script, "fi")
	if debianCheckEnd == -1 {
		return script
	}

	return script[debianCheckEnd+len("fi"):]
}

func buildYamlScript(header []byte) []byte {
	// TODO: remove after `preinstall` gets implemented
	script := removeDebianCheck(string(header)) + "\n"
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

	return []byte(script)
}

func RepairPacscript(pkg *types.Pacscript) error {
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

func (rp rawPacscript) toPacscript() types.Pacscript {
	out := types.Pacscript{
		Name:                 rp.Name,
		PackageName:          rp.PackageName,
		Maintainer:           rp.Maintainer,
		Description:          rp.Description,
		URL:                  rp.URL,
		RuntimeDependencies:  strings.Fields(rp.RuntimeDependencies),
		BuildDependencies:    strings.Fields(rp.BuildDependencies),
		Hash:                 rp.Hash,
		Breaks:               strings.Fields(rp.Breaks),
		Gives:                rp.Gives,
		Replace:              strings.Fields(rp.Replace),
		OptionalDependencies: rp.OptionalDependencies,
		PPA:                  rp.PPA,
		PacstallDependencies: rp.PacstallDependencies,
		Patch:                rp.Patch,
		RequiredBy:           make([]string, 0),
		Version:              rp.Version,
	}

	if out.PackageName == "" {
		out.PackageName = out.Name
	}

	return out
}

func computeRequiredBy(pkgs []*types.Pacscript) []*types.Pacscript {
	pickBeforeColon := func(arr []string) []string {
		out := make([]string, len(arr))
		for _, it := range arr {
			out = append(out, strings.Split(it, ":")[0])
		}
		return out
	}

	for _, pkg := range pkgs {
		pkg.RequiredBy = make([]string, 0)
		for _, otherPkg := range pkgs {

			for _, dependency := range pickBeforeColon(otherPkg.PacstallDependencies) {
				if dependency == pkg.Name {
					pkg.RequiredBy = append(pkg.RequiredBy, otherPkg.Name)
				}
			}
		}
	}

	return pkgs
}
