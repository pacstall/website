package ssr

import "regexp"

type IndexTemplateData struct {
	Title       string
	Description string
}

type IndexTemplateMaker func(path string, groups []string) IndexTemplateData

type SSRTemplate struct {
	PathMatcher   *regexp.Regexp
	TemplateMaker IndexTemplateMaker
}

var templates []SSRTemplate = make([]SSRTemplate, 0)

func AddTemplate(pathMatcher *regexp.Regexp, templateMaker IndexTemplateMaker) {
	templates = append(templates, SSRTemplate{
		PathMatcher:   pathMatcher,
		TemplateMaker: templateMaker,
	})
}

func GetTemplateForPath(path string) IndexTemplateData {
	for _, template := range templates {
		groups := template.PathMatcher.FindStringSubmatch(path)
		if len(groups) > 0 {
			return template.TemplateMaker(path, groups)
		}
	}

	return GenerateDefaultIndexTemplateData()
}

func GenerateDefaultIndexTemplateData() IndexTemplateData {
	return IndexTemplateData{
		Title:       "Pacstall - The AUR for Ubuntu",
		Description: "Pacstall automates downloading source packages, installing dependencies, and installing, in Ubuntu",
	}
}
