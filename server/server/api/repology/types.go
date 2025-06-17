package repology_api

import (
	"fmt"
	"strings"

	"pacstall.dev/webserver/types/pac"
)

type maintainerDetails struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type repologyPackage struct {
	Name              string                 `json:"name"`
	VisibleName       string                 `json:"visibleName"`
	Description       string                 `json:"description"`
	Maintainer        []maintainerDetails    `json:"maintainer"`
	Version           string                 `json:"version"`
	URL               []pac.ArchDistroString `json:"url"`
	RecipeURL         string                 `json:"recipeUrl"`
	PackageDetailsURL string                 `json:"packageDetailsUrl"`
	Type              string                 `json:"type"`
}

func newRepologyPackage(p *pac.Script) repologyPackage {
	return repologyPackage{
		Name:              p.PackageName,
		VisibleName:       p.PrettyName,
		Description:       p.Description,
		Maintainer:        getMaintainers(p),
		Version:           p.SourceVersion,
		URL:               p.Source,
		Type:              string(p.Type()),
		RecipeURL:         fmt.Sprintf("https://raw.githubusercontent.com/pacstall/pacstall-programs/master/packages/%s/%s.pacscript", p.PackageBase, p.PackageBase),
		PackageDetailsURL: fmt.Sprintf("https://pacstall.dev/packages/%s", p.PackageName),
	}
}

func getMaintainers(p *pac.Script) []maintainerDetails {
	maintainers := make([]maintainerDetails, 0, len(p.Maintainers))

	for _, m := range p.Maintainers {
		var name, email string
		if i := strings.Index(m, "<"); i != -1 && strings.HasSuffix(m, ">") {
			name = strings.TrimSpace(m[:i])
			email = strings.TrimSpace(m[i+1 : len(m)-1])
		} else {
			name = strings.TrimSpace(m)
		}
		maintainers = append(maintainers, maintainerDetails{
			Name:  &name,
			Email: &email,
		})
	}

	return maintainers
}
