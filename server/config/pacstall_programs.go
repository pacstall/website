package config

import (
	"path/filepath"
	"time"

	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/config"
)

var PacstallPrograms = config.PacstallProgramsConfig{}

func setPacstallPrograms(conf tomlPacstallProgramsConfig) {
	tempDir, err := filepath.Abs(conf.TempDir)
	if err != nil {
		log.Fatal("Could not parse file '%s'\n%v", *configPath, err)
	}

	PacstallPrograms = config.PacstallProgramsConfig{
		Path:           "programs",
		URL:            conf.URL,
		TempDir:        tempDir,
		UpdateInterval: time.Duration(conf.UpdateInterval) * time.Second,
		MaxOpenFiles:   conf.MaxOpenFiles,
	}
}
