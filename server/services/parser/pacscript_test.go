package parser

import (
	"testing"

	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/utils/expect"
)

func Test_buildCustomFormatScript(t *testing.T) {
	actual := buildCustomFormatScript("placeholder")
	expected := `placeholder
echo ''

if [[ "$(declare -F -p pkgname)" ]]; then
	pkgname=$(pkgname)
fi

if [[ "$(declare -F -p pkgdesc)" ]]; then
	pkgdesc=$(pkgdesc)
fi

if [[ "$(declare -F -p gives)" ]]; then
	gives=$(gives)
fi

if [[ "$(declare -F -p hash)" ]]; then
	hash=$(hash)
fi

if [[ "$(declare -F -p pkgver)" ]]; then
	pkgver=$(pkgver)
fi

jo -p -- -s pkgname="$pkgname" -s pkgdesc="$pkgdesc" -s gives="$gives" -s hash="$hash" -s pkgver="$pkgver" source=$(jo -a ${source[@]}) arch=$(jo -a ${arch[@]}) maintainer=$(jo -a ${maintainer[@]}) depends=$(jo -a ${depends[@]}) conflicts=$(jo -a ${conflicts[@]}) breaks=$(jo -a ${breaks[@]}) replaces=$(jo -a ${replaces[@]}) makedepends=$(jo -a ${makedepends[@]}) optdepends=$(jo -a ${optdepends[@]}) pacdeps=$(jo -a ${pacdeps[@]}) patch=$(jo -a ${patch[@]}) ppa=$(jo -a ${ppa[@]}) repology=$(jo -a ${repology[@]}) `

	expect.Equals(t, "json build script", expected, actual)
}

func Test_computeRequiredBy(t *testing.T) {
	data := []*pac.Script{
		{PackageName: "package_A", PacstallDependencies: []string{"package_B", "package_C"}},
		{PackageName: "package_B", PacstallDependencies: []string{"package_C"}},
		{PackageName: "package_C", PacstallDependencies: []string{}},
		{PackageName: "package_D", PacstallDependencies: []string{"package_A"}},
	}

	packageC := data[2]

	computeRequiredBy(packageC, data)

	actual := packageC.RequiredBy
	expected := []string{"package_A", "package_B"}

	expect.Equals(t, "required by", expected, actual)
}
