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

	script += "echo ''\n"
	for _, bashName := range pacsh.PacscriptVars {
		// If the variable is a function, then we replace it with the output of the function
		script += fmt.Sprintf(`
if [[ "$(declare -F -p %v)" ]]; then
	%v=$(%v)
fi
`, bashName, bashName, bashName)
	}

	script = script + "\njo -p -- "

	for _, bashName := range pacsh.PacscriptVars {
		script += fmt.Sprintf("-s %v=\"$%v\" ", bashName, bashName)
	}

	// `jo` json utility coerces all types to the closest matching.
	// So, for example, if a maintainer has the name "1234 me@example.com" it will
	// parse it as [1234, "me@example.com"], basically [int, string] array.
	// We need it to be a pure string array ["1234", "me@example.com"].
	// By adding an underscore prefix we mitigate this issue of coercing to number but
	// we have to remove it later in the json parsing
	for _, bashName := range pacsh.PacscriptArrays {
		script += fmt.Sprintf("%v=$(jo -a _${%v[@]}) ", bashName, bashName)
	}

	return []byte(script)
}

func computeRequiredBy(script *pac.Script, scripts []*pac.Script) {
	pickBeforeColon := func(it *array.Iterator[string]) string {
		return strings.Split(it.Value, ": ")[0]
	}

	script.RequiredBy = make([]string, 0)
	for _, otherScript := range scripts {
		otherScriptDependencies := array.Map(otherScript.PacstallDependencies, pickBeforeColon)
		if array.Contains(otherScriptDependencies, array.Is(script.PackageName)) {
			script.RequiredBy = append(script.RequiredBy, otherScript.PackageName)
		}
	}
}
