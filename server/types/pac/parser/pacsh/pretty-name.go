package pacsh

import (
	"strings"

	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

var pacTypes = map[string]string{
	"-deb": "Debian Native",
	"-git": "Source Code",
	"-bin": "Precompiled",
	"-app": "AppImage",
}

func getPrettyName(p pac.Script) string {
	name := ""

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

	return out
}
