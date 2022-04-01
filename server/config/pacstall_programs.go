package config

import "time"

type PacstallProgramsConfig struct {
	Path           string
	TempDir        string
	UpdateInterval time.Duration
	MaxOpenFiles   uint8
}

var PacstallPrograms = PacstallProgramsConfig{}

func setPacstallPrograms(conf tomlPacstallProgramsConfig) {
	PacstallPrograms = PacstallProgramsConfig{
		Path:           conf.Path,
		TempDir:        conf.TempDir,
		UpdateInterval: time.Duration(conf.UpdateInterval) * time.Second,
		MaxOpenFiles:   conf.MaxOpenFiles,
	}
}
