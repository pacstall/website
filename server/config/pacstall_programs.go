package config

import (
	"path/filepath"
	"time"

	"pacstall.dev/webserver/log"
)

type PacstallProgramsConfig struct {
	Path           string
	TempDir        string
	UpdateInterval time.Duration
	MaxOpenFiles   uint8
}

var PacstallPrograms = PacstallProgramsConfig{}

func setPacstallPrograms(conf tomlPacstallProgramsConfig) {
	path, err := filepath.Abs(conf.Path)

	if err != nil {
		log.Error.Fatalf("Could not parse file '%s'\n%v", *configPath, err)
	}

	tempDir, err := filepath.Abs(conf.TempDir)
	if err != nil {
		log.Error.Fatalf("Could not parse file '%s'\n%v", *configPath, err)
	}

	PacstallPrograms = PacstallProgramsConfig{
		Path:           path,
		TempDir:        tempDir,
		UpdateInterval: time.Duration(conf.UpdateInterval) * time.Second,
		MaxOpenFiles:   conf.MaxOpenFiles,
	}
}
