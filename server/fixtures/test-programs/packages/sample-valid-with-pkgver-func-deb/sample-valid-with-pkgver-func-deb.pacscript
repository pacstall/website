pkgname="sample-valid-deb"
gives="sample-valid"
repology=("project: sample-valid")
source=("https://example.com")
pkgdesc="Sample description"
hash="10101010"
arch=('amd64')
makedepends=("go" "gcc")
optdepends=("opt1" "opt2")
pacdeps=("pacdep1" "pacdep2")
depends=("dep1" "dep2")
ppa=("ppa1" "ppa2")
patch=("patch1" "patch2")
provides=("provides1" "provides2")
incompatible=("incompatible1" "incompatible2")
maintainer=("pacstall <test@pacstall.dev>")
breaks=("breaks1" "breaks2")
conflicts=("conflicts1" "conflicts2")
replaces=("replaces1" "replaces2")

pkgver() {
    echo "1.2.3"
}