package config

import (
	"path/filepath"

	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/config"
)

var TCPServer = config.TCPServerConfig{}

func setTCPServer(conf tomlTCPServerConfig) {
	publicDir, err := filepath.Abs(conf.PublicDir)
	if err != nil {
		log.Fatal("Could not parse file '%s'\n%v", *configPath, err)
	}

	TCPServer.Port = conf.Port
	TCPServer.PublicDir = publicDir
}
