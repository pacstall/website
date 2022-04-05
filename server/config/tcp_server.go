package config

import (
	"path/filepath"

	"pacstall.dev/webserver/log"
)

type TCPServerConfig struct {
	Port      uint16
	PublicDir string
}

var TCPServer = TCPServerConfig{}

func setTCPServer(conf tomlTCPServerConfig) {
	publicDir, err := filepath.Abs(conf.PublicDir)
	if err != nil {
		log.Error.Fatalf("Could not parse file '%s'\n%v", *configPath, err)
	}

	TCPServer.Port = conf.Port
	TCPServer.PublicDir = publicDir
}
