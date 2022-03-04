package pacscript

import (
	"time"

	"pacstall.dev/website/types"
)

var lastModified time.Time
var loadedPackages []*types.PackageInfo
