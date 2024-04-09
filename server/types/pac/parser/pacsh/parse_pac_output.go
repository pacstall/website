package pacsh

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/types/pac"
)

var ParsePacOutput = parseOutput
var PacscriptVars []string = []string{"pkgname", "pkgdesc", "gives", "hash", "pkgver"}
var PacscriptArrays []string = []string{"source", "arch", "maintainer", "depends", "conflicts", "breaks", "replaces", "makedepends", "optdepends", "pacdeps", "patch", "ppa", "repology"}

type pacscriptJsonStructure struct {
	Pkgname     string   `json:"pkgname"`
	Pkgdesc     string   `json:"pkgdesc"`
	Gives       *string  `json:"gives"`
	Hash        *string  `json:"hash"`
	Pkgver      *string  `json:"pkgver"`
	Source      []string `json:"source"`
	Maintainer  []string `json:"maintainer"`
	Depends     []string `json:"depends"`
	Conflicts   []string `json:"conflicts"`
	Arch        []string `json:"arch"`
	Breaks      []string `json:"breaks"`
	Replaces    []string `json:"replaces"`
	Makedepends []string `json:"makedepends"`
	Optdepends  []string `json:"optdepends"`
	Pacdeps     []string `json:"pacdeps"`
	Patch       []string `json:"patch"`
	Ppa         []string `json:"ppa"`
	Repology    []string `json:"repology"`
}

var _GIT_VERSION = "git"
var _EMPTI_STR = ""

func parseOutput(data []byte) (out pac.Script, err error) {
	var parsedContent pacscriptJsonStructure
	err = json.Unmarshal(data, &parsedContent)
	if err != nil {
		return out, errorx.IllegalFormat.New("failed to deserialize json content '%v'. err: %v", string(data), err)
	}

	if parsedContent.Pkgver == nil {
		if strings.HasSuffix(parsedContent.Pkgname, "-git") {
			parsedContent.Pkgver = &_GIT_VERSION
		} else {
			return out, errorx.IllegalArgument.New("expected version to be non-empty")
		}
	}

	if parsedContent.Gives == nil {
		parsedContent.Gives = &_EMPTI_STR
	}

	out = pac.Script{
		PackageName:          parsedContent.Pkgname,
		Maintainers:          parseMaintainers(parsedContent.Maintainer),
		Description:          parsedContent.Pkgdesc,
		Source:               parsedContent.Source,
		Gives:                *parsedContent.Gives,
		Hash:                 parsedContent.Hash,
		Version:              *parsedContent.Pkgver,
		RuntimeDependencies:  parsedContent.Depends,
		Conflicts:            parsedContent.Conflicts,
		Breaks:               parsedContent.Breaks,
		Replaces:             parsedContent.Replaces,
		BuildDependencies:    parsedContent.Makedepends,
		OptionalDependencies: parsedContent.Optdepends,
		PacstallDependencies: parsedContent.Pacdeps,
		PPA:                  parsedContent.Ppa,
		Patch:                parsedContent.Patch,
		RequiredBy:           make([]string, 0),
		Repology:             parsedContent.Repology,
		LatestVersion:        nil,
		UpdateStatus:         pac.UpdateStatus.Unknown,
	}

	// Fixes up repology project split
	if len(out.Repology) > 1 && out.Repology[0] == "project:" {
		repology := []string{}
		repology = append(repology, fmt.Sprintf("project: %v", out.Repology[1]))
		repology = append(repology, out.Repology[2:]...)
		out.Repology = repology
	}

	if out.Hash != nil && len(*out.Hash) == 0 {
		out.Hash = nil
	}

	out.PrettyName = getPrettyName(out)

	return
}

func parseMaintainers(maintainers []string) []string {
	maintainersSplitByLA := strings.Split(strings.Join(maintainers, " "), ">")

	out := []string{}
	for _, maintainer := range maintainersSplitByLA {
		if len(maintainer) == 0 {
			continue
		}
		out = append(out, strings.TrimSpace(maintainer)+">")
	}

	return out
}
