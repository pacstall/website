package ssr

import (
	"fmt"
	"regexp"

	r "pacstall.dev/webserver/ssr"
	"pacstall.dev/webserver/store/pacstore"
)

func registerPacscriptSSRData() {
	r.AddTemplate(
		regexp.MustCompile(`^/packages/([a-zA-Z0-9-]+)`),
		func(path string, groups []string) r.IndexTemplateData {
			name := groups[1]

			pkg, err := pacstore.GetAll().FindByName(name)
			if err != nil {
				return r.GenerateDefaultIndexTemplateData()
			}

			return r.IndexTemplateData{
				Title:       fmt.Sprintf("%s - Pacstall", pkg.Name),
				Description: pkg.Description,
			}
		},
	)
}
