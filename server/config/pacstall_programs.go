package config

import (
	"path/filepath"
	"time"

	"pacstall.dev/webserver/log"
)

type PacstallProgramsConfig struct {
	Path           string
	URL            string
	TempDir        string
	UpdateInterval time.Duration
	MaxOpenFiles   uint8
}

var PacstallPrograms = PacstallProgramsConfig{}

func setPacstallPrograms(conf tomlPacstallProgramsConfig) {
	tempDir, err := filepath.Abs(conf.TempDir)
	if err != nil {
		log.Error.Fatalf("Could not parse file '%s'\n%v", *configPath, err)
	}

	PacstallPrograms = PacstallProgramsConfig{
		Path:           "programs",
		URL:            conf.URL,
		TempDir:        tempDir,
		UpdateInterval: time.Duration(conf.UpdateInterval) * time.Second,
		MaxOpenFiles:   conf.MaxOpenFiles,
	}
}
