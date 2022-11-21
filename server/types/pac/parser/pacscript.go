package parser

import (
	"fmt"
	"strings"

	"pacstall.dev/webserver/types/pac/parser/pacsh"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

func removeDebianCheck(script string) string {
	if !strings.Contains(script, "/etc/os-release)\" == \"debian\"") {
		return script
	}

	if strings.Index(script, "if") != 0 {
		return script
	}

	debianCheckEnd := strings.Index(script, "fi")
	if debianCheckEnd == -1 {
		return script
	}

	return script[debianCheckEnd+len("fi"):]
}

func buildCustomFormatScript(header []byte) []byte {
	// TODO: remove after `preinstall` gets implemented
	script := removeDebianCheck(string(header)) + "\n"

	categoryToken := `++++`
	subcategoryToken := `+  +++`

	script = script + "echo ''\n"

	for _, bashName := range pacsh.PacstallCVars {
		script += fmt.Sprintf("echo \"%s $%v\"", categoryToken, bashName) + "\n"
	}

	for _, bashName := range pacsh.PacstallCArrays {
		script += fmt.Sprintf("echo \"%s $%v\"", categoryToken, bashName) + "\n"
	}

	mapsPartialScript := make([]string, 0)
	for _, bashName := range pacsh.PacstallCMaps {
		partial := "echo " + categoryToken + "\n"

		partial += fmt.Sprintf("printf '%s%%s\\n' \"${%v[@]}\"\n", subcategoryToken, bashName)
		mapsPartialScript = append(mapsPartialScript, partial)
	}

	script += strings.Join(mapsPartialScript, "\n")

	return []byte(script)
}

func computeRequiredBy(script pac.Script, scripts list.List[*pac.Script]) *pac.Script {
	pickBeforeColon := func(line string) string {
		return strings.Split(line, ": ")[0]
	}

	script.RequiredBy = list.Map(
		scripts.Filter(func(s *pac.Script) bool {
			return list.From(s.PacstallDependencies).Map(pickBeforeColon).Contains(list.Is(script.Name))
		}), func(_ int, s *pac.Script) string {
			return s.Name
		})

	return &script
}
