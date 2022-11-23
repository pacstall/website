package config

import (
	"fmt"
	"strconv"
	"time"

	"pacstall.dev/webserver/config/build"
)

var Production = false

var UpdateInterval = time.Duration(toInt(build.UpdateInterval)) * time.Second
var TempDir = build.TempDir
var MaxOpenFiles = toInt(build.MaxOpenFiles)
var GitURL = build.GitURL
var GitClonePath = build.GitClonePath

var Port = toInt(build.Port)
var PublicDir = build.PublicDir

var Version = build.PublicDir

func toInt(str string) int {
	die := func (err error, message string, args ...any) {
		if err != nil {
			fmt.Printf(message, args...)
			panic(err)
		}
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		die(err, "could not convert '%s' to int\n", str)
	}

	return num
}