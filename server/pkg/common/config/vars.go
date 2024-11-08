package config

import (
	"pacstall.dev/webserver/pkg/common/config/build"
	"pacstall.dev/webserver/pkg/common/env"
)

var Production = build.Production == "true" || build.Production == "1"
var Version = build.Version

var LogLevel = env.GetEnvStringOrDefault("PACSTALL_LOG_LEVEL", "debug")
