package ssr

import (
	"regexp"

	r "pacstall.dev/webserver/server/ssr"
)

func registerPacscriptListSSRData() {
	r.AddTemplate(
		regexp.MustCompile(`^/packages[/#]?(\?[a-zA-Z0-9-=&]*)?$`),
		func(path string, groups []string) r.IndexTemplateData {
			return r.IndexTemplateData{
				Title:       "Packages | Pacstall",
				Description: r.GenerateDefaultIndexTemplateData().Description,
				Html:        r.GenerateDefaultIndexTemplateData().Html,
			}
		},
	)
}
