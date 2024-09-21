package pacsh

import (
	"strings"

	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/pac"
)

func getPrettyName(p *pac.Script) string {
	name := ""

	if name == "" {
		name = p.PackageName
	}

	for suffix := range types.PackageTypeSuffixToPackageTypeName {
		if strings.HasSuffix(name, string(suffix)) {
			name = name[0 : len(name)-len(suffix)]
		}
	}

	return titleCase(name)
}

func titleCase(s string) string {
	title := ""
	words := strings.Split(s, "-")

	for _, word := range words {
		if title != "" {
			title += " "
		}

		if len(word) == 0 {
			continue
		}
		title += strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}

	return title
}
