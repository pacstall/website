package ssr

import "regexp"

type IndexTemplateData struct {
	Title       string
	Description string
	Html		string
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
		Html: `
			<h1>Pacstall - The AUR for Ubuntu<h1>
			<h2>Homepage<h2>
			<nav>
				<ul>
					<li><a href="/packages">Browse Packages</a></li>
					<li><a href="/privacy">Privacy Policy</a></li>
				</ul>
			</nav>
			<main>
				<h3>Pacstall automates downloading source packages, installing dependencies, and installing, in Ubuntu</h3>

				<article>
					<a href="https://github.com/pacstall/pacstall/wiki/How-to-contribute">Contribute</a>
				</article>
				<article>
					<a href="https://github.com/pacstall/pacstall/wiki/Pacscript-101">Become a Package Maintainer</a>
				</article>
			</main>
		`,
	}
}
