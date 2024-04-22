package pacsh

import (
	"encoding/json"
	"strings"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
)

var ParsePacOutput = parseOutput
var PacscriptVars []string = []string{"pkgname", "pkgdesc", "gives", "hash", "pkgver"}
var PacscriptArrays []string = []string{"source", "arch", "maintainer", "depends", "conflicts", "breaks", "replaces", "makedepends", "optdepends", "pacdeps", "patch", "ppa", "repology"}

type Stringable struct {
	Data string
}

func (s *Stringable) UnmarshalJSON(data []byte) error {
	s.Data = string(data)
	s.Data = strings.ReplaceAll(s.Data, "\"", "")
	return nil
}

func (s *Stringable) String() string {
	return s.Data
}

type StringableArrayWithComments struct {
	Data []Stringable
}

func (s StringableArrayWithComments) toStringArray() []string {
	return array.SwitchMapPtr(s.Data, func(it *array.PtrIterator[Stringable]) string {
		val := *it.Value
		return val.String()
	})
}

func (s *StringableArrayWithComments) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &s.Data)
	if err != nil {
		return err
	}

	out := make([]string, 0)
	item := ""
	hasComments := false
	for _, word := range s.Data {
		strWord := word.String()
		wordHasComment := strings.HasSuffix(strWord, ":")

		if wordHasComment {
			hasComments = true
			if item != "" {
				out = append(out, strings.TrimSpace(item))
			}
			item = strWord
		} else if !hasComments {
			out = append(out, strWord)
		} else {
			item += " " + strWord
		}

	}

	if item != "" {
		out = append(out, strings.TrimSpace(item))
	}

	s.Data = array.SwitchMap(out, func(it *array.Iterator[string]) Stringable {
		return Stringable{it.Value}
	})

	return nil
}

type pacscriptJsonStructure struct {
	Pkgname     string                      `json:"pkgname"`
	Pkgdesc     string                      `json:"pkgdesc"`
	Gives       *string                     `json:"gives"`
	Hash        *string                     `json:"hash"`
	Pkgver      *string                     `json:"pkgver"`
	Source      StringableArrayWithComments `json:"source"`
	Maintainer  StringableArrayWithComments `json:"maintainer"`
	Depends     StringableArrayWithComments `json:"depends"`
	Conflicts   StringableArrayWithComments `json:"conflicts"`
	Arch        StringableArrayWithComments `json:"arch"`
	Breaks      StringableArrayWithComments `json:"breaks"`
	Replaces    StringableArrayWithComments `json:"replaces"`
	Makedepends StringableArrayWithComments `json:"makedepends"`
	Optdepends  StringableArrayWithComments `json:"optdepends"`
	Pacdeps     StringableArrayWithComments `json:"pacdeps"`
	Patch       StringableArrayWithComments `json:"patch"`
	Ppa         StringableArrayWithComments `json:"ppa"`
	Repology    StringableArrayWithComments `json:"repology"`
}

var _GIT_VERSION = "git"
var _EMPTI_STR = ""

func parseOutput(data []byte) (out pac.Script, err error) {
	// remove prefixes if any
	runeIndex := strings.IndexRune(string(data), '{')
	if runeIndex >= 0 {
		str := string(data)
		str = str[runeIndex:]
		data = []byte(str)
	}

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
		Maintainers:          parseMaintainers(parsedContent.Maintainer.toStringArray()),
		Description:          parsedContent.Pkgdesc,
		Source:               parsedContent.Source.toStringArray(),
		Gives:                *parsedContent.Gives,
		Hash:                 parsedContent.Hash,
		Version:              *parsedContent.Pkgver,
		RuntimeDependencies:  parsedContent.Depends.toStringArray(),
		BuildDependencies:    parsedContent.Makedepends.toStringArray(),
		OptionalDependencies: parsedContent.Optdepends.toStringArray(),
		Conflicts:            parsedContent.Conflicts.toStringArray(),
		Replaces:             parsedContent.Replaces.toStringArray(),
		Breaks:               parsedContent.Breaks.toStringArray(),
		PacstallDependencies: parsedContent.Pacdeps.toStringArray(),
		PPA:                  parsedContent.Ppa.toStringArray(),
		Patch:                parsedContent.Patch.toStringArray(),
		RequiredBy:           make([]string, 0),
		Repology:             parsedContent.Repology.toStringArray(),
		LatestVersion:        nil,
		UpdateStatus:         pac.UpdateStatus.Unknown,
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
