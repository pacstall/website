package repology_api

import (
	"fmt"
	"strings"

	"pacstall.dev/webserver/consts"
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
	URL               *string           `json:"url"`
	RecipeURL         string            `json:"recipeUrl"`
	PackageDetailsURL string            `json:"packageDetailsUrl"`
	Type              string            `json:"type"`
	Patches           []string          `json:"patches"`
}

func newRepologyPackage(p *pac.Script) repologyPackage {
	var source *string = nil
	if len(p.Source) > 0 {
		source = &p.Source[0]
	}

	return repologyPackage{
		Name:              p.PackageName,
		VisibleName:       p.PrettyName,
		Description:       p.Description,
		Maintainer:        getMaintainer(p),
		Version:           p.Version,
		URL:               source,
		Type:              getType(p),
		RecipeURL:         fmt.Sprintf("https://raw.githubusercontent.com/pacstall/pacstall-programs/master/packages/%s/%s.%s", p.PackageName, p.PackageName, consts.PACSCRIPT_FILE_EXTENSION),
		PackageDetailsURL: fmt.Sprintf("https://pacstall.dev/packages/%s", p.PackageName),
		Patches:           p.Patch,
	}
}

var pacTypes = map[string]string{
	"-deb": "Debian Native",
	"-git": "Source Code",
	"-bin": "Precompiled",
	"-app": "AppImage",
}

func getMaintainer(p *pac.Script) maintainerDetails {
	maintainer := ""
	if len(p.Maintainers) > 0 {
		maintainer = p.Maintainers[0]
	}

	if !strings.Contains(maintainer, "<") {
		return maintainerDetails{
			Name: &maintainer,
		}
	}

	parts := strings.Split(maintainer, "<")
	name := strings.TrimSpace(parts[0])
	email := strings.TrimSpace(strings.Replace(parts[1], ">", "", -1))

	return maintainerDetails{
		Name:  &name,
		Email: &email,
	}
}

func getType(p *pac.Script) string {
	for suffix, kind := range pacTypes {
		if strings.HasSuffix(p.PackageName, suffix) {
			return kind
		}
	}

	return pacTypes["-git"]
}
