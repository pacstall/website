package ssr

import (
	"regexp"
)

func (s *ServerSideRenderService) registerPacscriptListSSRData() {
	AddTemplate(
		regexp.MustCompile(`^/packages[/#]?(\?[a-zA-Z0-9-=&]*)?$`),
		func(path string, groups []string) IndexTemplateData {
			return IndexTemplateData{
				Title:       "Packages | Pacstall",
				Description: GenerateDefaultIndexTemplateData().Description,
				Html:        GenerateDefaultIndexTemplateData().Html,
			}
		},
	)
}
