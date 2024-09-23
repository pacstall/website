package config

import (
	"time"

	"pacstall.dev/webserver/pkg/common/env"
)

var PacSight = struct {
	Port int
}{}

func initPacSightEnv() {
	PacSight.Port = env.GetEnvIntOrDefault("PACSTALL_PACSIGHT_PORT", 8080)
}

var Repology = struct {
	RepologyUpdateInterval time.Duration
}{}

func initRepologyEnv() {
	Repology.RepologyUpdateInterval = time.Duration(env.GetEnvIntOrDefault("PACSTALL_REPOLOGY_UPDATE_INTERVAL", 60*60*6)) * time.Second
}

func Init() {
	initPacSightEnv()
	initRepologyEnv()
}
