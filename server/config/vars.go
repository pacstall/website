package config

import (
	"time"

	"pacstall.dev/webserver/config/build"
)

var Production = toBool(build.Production)

var UpdateInterval = time.Duration(toInt(build.UpdateInterval)) * time.Second
var TempDir = build.TempDir
var MaxOpenFiles = toInt(build.MaxOpenFiles)
var GitURL = build.GitURL
var GitClonePath = build.GitClonePath

var Port = toInt(build.Port)
var PublicDir = build.PublicDir

var Version = build.Version
