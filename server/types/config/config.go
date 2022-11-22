package config

import "time"

type PacstallProgramsConfig struct {
	Path           string
	URL            string
	TempDir        string
	UpdateInterval time.Duration
	MaxOpenFiles   uint8
}
