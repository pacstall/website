package config

import (
	"pacstall.dev/webserver/pkg/common/config/build"
)

var Production = build.Production == "true" || build.Production == "1"
var Version = build.Version
