package parser

import (
	"fmt"
	"strings"

	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/parser/pacsh"
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
		// If the variable is a function, then we replace it with the output of the function
		script += fmt.Sprintf(`
if [[ "$(declare -F -p %v)" ]]; then
	%v=$(%v)
fi
`, bashName, bashName, bashName)
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

func computeRequiredBy(script *pac.Script, scripts []*pac.Script) {
	pickBeforeColon := func(it *array.Iterator[string]) string {
		return strings.Split(it.Value, ": ")[0]
	}

	script.RequiredBy = make([]string, 0)
	for _, otherScript := range scripts {
		otherScriptDependencies := array.Map(otherScript.PacstallDependencies, pickBeforeColon)
		if array.Contains(otherScriptDependencies, array.Is(script.Name)) {
			script.RequiredBy = append(script.RequiredBy, otherScript.Name)
		}
	}
}
