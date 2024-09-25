package config

import (
	"time"

	"pacstall.dev/webserver/pkg/common/env"
)

var PacSight = struct {
	Port int
}{}

func initPacSightEnv() {
	PacSight.Port = env.GetEnvIntOrDefault("PACSTALL_PACSIGHT_PORT", 3301)
}

var Repology = struct {
	RepologyUpdateInterval time.Duration
	CachePath              string
	MaxOpenFiles           int
}{}

func initRepologyEnv() {
	Repology.RepologyUpdateInterval = time.Duration(env.GetEnvIntOrDefault("PACSTALL_REPOLOGY_UPDATE_INTERVAL", 60*60*6)) * time.Second
	Repology.CachePath = env.GetEnvStringOrDefault("PACSTALL_REPOLOGY_CACHE_PATH", "./repology_cache")
	Repology.MaxOpenFiles = env.GetEnvIntOrDefault("PACSTALL_REPOLOGY_MAX_OPEN_FILES", 10)
}

func Init() {
	initPacSightEnv()
	initRepologyEnv()
}
