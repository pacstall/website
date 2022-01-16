package pacpkgs

import "pacstall.dev/website/types"

func GetPackages() []types.PackageInfo {
	return loadedPackages
}
