name="foo"
pkgname="foo" # if this is a -git, -bin, etc. package, this would be the package name without extension
version="1.0"
url="https://github.com/Henryws/foo/archive/refs/tags/${version}.zip"
build_depends="vim gcc"
depends="neofetch plasma"
breaks="${pkgname} ${pkgname}-git ${pkgname}-bin ${pkgname}-app ${pkgname}-deb"
replace="${pkgname} alacritty"
#gives="${pkgname}"
description="foo is the ultimate program capable of foo and bar!"
hash="2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
optdepends=("bar: not foo"
"alacritty: a blazing fast terminal")
ppa=("graphics-drivers/ppa"
"webupd8team/y-ppa-manager")
maintainer="Mr. Person <mr.person@protonmail.com>"
pacdeps=("dmenu" "tuner")
patch=("https://dwm.suckless.org/patches/anybar/dwm-anybar-20200810-bb2e722.diff")