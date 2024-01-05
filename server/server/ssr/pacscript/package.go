package ssr

import (
	"fmt"
	"regexp"

	"pacstall.dev/webserver/consts"
	r "pacstall.dev/webserver/server/ssr"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/pacstore"
)

func registerPacscriptSSRData() {
	r.AddTemplate(
		regexp.MustCompile(`^/packages/([a-zA-Z0-9-]+)`),
		func(path string, groups []string) r.IndexTemplateData {
			name := groups[1]

			pkg, err := array.FindBy(pacstore.GetAll(), func(s *pac.Script) bool {
				return s.Name == name
			})

			if err != nil {
				return r.GenerateDefaultIndexTemplateData()
			}

			return r.IndexTemplateData{
				Title:       fmt.Sprintf("%s - Pacstall", pkg.Name),
				Description: pkg.Description,
				Html: fmt.Sprintf(`
				<h1>Pacstall - The AUR for Ubuntu<h1>
				<h2>%s<h2>
				<nav>
					<ul>
						<li><a href="/packages">Browse Packages</a></li>
						<li><a href="/privacy">Privacy Policy</a></li>
					</ul>
				</nav>
				<main>
					<article>
						<h3>Package: %s</h3>
						<h3>Description: %s</h3>
						<h3>Maintainer: %s</h3>
						<h3>Version: %s</h3>
						<h3><a href="%s">URL</a></h3>
						<h3><a href="https://github.com/pacstall/pacstall-programs/blob/master/packages/%s/%s.%s">Source</a></h3>
					</article>

					<button>Install now!</button>
					<p>Find similar packages <a href="/packages?page=0&size=25&sortBy=default&sort=asc&filter=%s&filterBy=name">here</a>.</p>
				</main>
			`, pkg.Name, pkg.Name, pkg.Description, pkg.Maintainer, pkg.Version, pkg.URL, pkg.Name, pkg.Name, consts.PACSCRIPT_FILE_EXTENSION, pkg.PackageName),
			}
		},
	)
}
