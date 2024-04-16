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
		Maintainers:          parseMaintainers(removePrefixFromArray(parsedContent.Maintainer)),
		Description:          parsedContent.Pkgdesc,
		Source:               removePrefixFromArray(parsedContent.Source),
		Gives:                *parsedContent.Gives,
		Hash:                 parsedContent.Hash,
		Version:              *parsedContent.Pkgver,
		RuntimeDependencies:  removePrefixFromArray(parsedContent.Depends),
		BuildDependencies:    removePrefixFromArray(parsedContent.Makedepends),
		OptionalDependencies: removePrefixFromArray(parsedContent.Optdepends),
		Conflicts:            removePrefixFromArray(parsedContent.Conflicts),
		Replaces:             removePrefixFromArray(parsedContent.Replaces),
		Breaks:               removePrefixFromArray(parsedContent.Breaks),
		PacstallDependencies: removePrefixFromArray(parsedContent.Pacdeps),
		PPA:                  removePrefixFromArray(parsedContent.Ppa),
		Patch:                removePrefixFromArray(parsedContent.Patch),
		RequiredBy:           make([]string, 0),
		Repology:             removePrefixFromArray(parsedContent.Repology),
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

func removePrefix(word string) string {
	if strings.HasPrefix(word, "_") {
		return word[1:]
	}

	return word
}

func removePrefixFromArray(words []string) []string {
	if len(words) == 0 {
		return words
	}

	words[0] = removePrefix(words[0])
	if len(words) == 1 && len(words[0]) == 0 {
		return []string{}
	}

	return words
}
