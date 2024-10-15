package types

type PackageTypeName string

const (
	PACKAGE_TYPE_DEB PackageTypeName = "Debian Native"
	PACKAGE_TYPE_GIT PackageTypeName = "Source Code"
	PACKAGE_TYPE_BIN PackageTypeName = "Precompiled"
	PACKAGE_TYPE_APP PackageTypeName = "AppImage"
)

type PackageTypeSuffix string

const (
	PACKAGE_TYPE_SUFFIX_DEB PackageTypeSuffix = "-deb"
	PACKAGE_TYPE_SUFFIX_GIT PackageTypeSuffix = "-git"
	PACKAGE_TYPE_SUFFIX_BIN PackageTypeSuffix = "-bin"
	PACKAGE_TYPE_SUFFIX_APP PackageTypeSuffix = "-app"
)

var PackageTypeSuffixToPackageTypeName = map[PackageTypeSuffix]PackageTypeName{
	PACKAGE_TYPE_SUFFIX_DEB: PACKAGE_TYPE_DEB,
	PACKAGE_TYPE_SUFFIX_GIT: PACKAGE_TYPE_GIT,
	PACKAGE_TYPE_SUFFIX_BIN: PACKAGE_TYPE_BIN,
	PACKAGE_TYPE_SUFFIX_APP: PACKAGE_TYPE_APP,
}
