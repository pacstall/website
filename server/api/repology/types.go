package repology_api

import (
	"fmt"
	"strings"

	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

type maintainerDetails struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type repologyPackage struct {
	Name              string            `json:"name"`
	VisibleName       string            `json:"visibleName"`
	Description       string            `json:"description"`
	Maintainer        maintainerDetails `json:"maintainer"`
	Version           string            `json:"version"`
	URL               string            `json:"url"`
	RecipeURL         string            `json:"recipeUrl"`
	PackageDetailsURL string            `json:"packageDetailsUrl"`
	Type              string            `json:"type"`
	Patches           []string          `json:"patches"`
}

func newRepologyPackage(p pac.Script) repologyPackage {
	return repologyPackage{
		Name:              p.Name,
		VisibleName:       getPrettyName(p),
		Description:       p.Description,
		Maintainer:        getMaintainer(p),
		Version:           p.Version,
		URL:               p.URL,
		Type:              getType(p),
		RecipeURL:         fmt.Sprintf("https://raw.githubusercontent.com/pacstall/pacstall-programs/master/packages/%s/%s.pacscript", p.Name, p.Name),
		PackageDetailsURL: fmt.Sprintf("https://pacstall.dev/packages/%s", p.Name),
		Patches:           p.Patch,
	}
}

var pacTypes = map[string]string{
	"-deb": "Debian Native",
	"-git": "Source Code",
	"-bin": "Precompiled",
	"-app": "AppImage",
}

func getPrettyName(p pac.Script) string {
	name := p.PrettyName

	if name == "" {
		name = p.Name
	}

	for suffix := range pacTypes {
		if strings.HasSuffix(name, suffix) {
			name = name[0 : len(name)-len(suffix)]
		}
	}

	return titleCase(name)
}

func titleCase(s string) string {
	out := list.Reduce(strings.Split(s, "-"), func(word string, acc string) string {
		if acc != "" {
			acc += " "
		}
		return acc + strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}, "")

	log.Warn.Printf("Title case: %s -> %s", s, out)

	return out
}

func getMaintainer(p pac.Script) maintainerDetails {
	if !strings.Contains(p.Maintainer, "<") {
		return maintainerDetails{
			Name: &p.Maintainer,
		}
	}

	parts := strings.Split(p.Maintainer, "<")
	name := strings.TrimSpace(parts[0])
	email := strings.TrimSpace(strings.Replace(parts[1], ">", "", -1))

	return maintainerDetails{
		Name:  &name,
		Email: &email,
	}
}

func getType(p pac.Script) string {
	for suffix, kind := range pacTypes {
		if strings.HasSuffix(p.Name, suffix) {
			return kind
		}
	}

	return pacTypes["-git"]
}
