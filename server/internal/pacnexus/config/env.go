package config

import (
	"time"

	"pacstall.dev/webserver/pkg/common/env"
)

var PacSight = struct {
	Port int
	Host string
}{}

func initPacSightEnv() {
	PacSight.Port = env.GetEnvIntOrDefault("PACSTALL_PACSIGHT_PORT", 3301)
	PacSight.Host = env.GetEnvStringOrDefault("PACSTALL_PACSIGHT_HOST", "localhost")
}

var PacNexus = struct {
	Port      int
	PublicDir string
}{}

func initPacNexusEnv() {
	PacNexus.Port = env.GetEnvIntOrDefault("PACSTALL_PACNEXUS_PORT", 3300)
	PacNexus.PublicDir = env.GetEnvStringOrDefault("PACSTALL_PACNEXUS_PUBLIC_DIR", "./public")
}

// Configuration for the discord integration
var Discord = struct {
	Token     string
	ChannelID string
	Enabled   bool
	Tags      string
}{}

func initDiscordEnv() {
	Discord.Enabled = env.GetEnvBoolOrDefault("PACSTALL_DISCORD_ENABLED", false)

	if !Discord.Enabled {
		return
	}

	Discord.Token = env.GetEnvString("PACSTALL_DISCORD_TOKEN")
	Discord.ChannelID = env.GetEnvString("PACSTALL_DISCORD_CHANNEL_ID")
	Discord.Tags = env.GetEnvString("PACSTALL_DISCORD_TAGS")
}

var PacstallPrograms = struct {
	Branch         string
	UpdateInterval time.Duration
	TempDir        string
	MaxOpenFiles   int
	GitURL         string
	GitClonePath   string
}{}

func initPacstallProgramsEnv() {
	PacstallPrograms.Branch = env.GetEnvStringOrDefault("PACSTALL_PROGRAMS_GIT_BRANCH", "master")
	PacstallPrograms.UpdateInterval = time.Duration(env.GetEnvIntOrDefault("PACSTALL_PROGRAMS_UPDATE_INTERVAL", 15*60)) * time.Second
	PacstallPrograms.TempDir = env.GetEnvStringOrDefault("PACSTALL_PROGRAMS_TEMP_DIR", "./tmp")
	PacstallPrograms.MaxOpenFiles = env.GetEnvIntOrDefault("PACSTALL_PROGRAMS_MAX_OPEN_FILES", 100)
	PacstallPrograms.GitURL = env.GetEnvStringOrDefault("PACSTALL_PROGRAMS_GIT_URL", "https://github.com/pacstall/pacstall-programs.git")
	PacstallPrograms.GitClonePath = env.GetEnvStringOrDefault("PACSTALL_PROGRAMS_GIT_CLONE_PATH", "./programs")
}

// Configuration for the database
var Database = struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}{}

func initDatabaseEnv() {
	Database.Host = env.GetEnvStringOrDefault("PACSTALL_DATABASE_HOST", "localhost")
	Database.Port = env.GetEnvIntOrDefault("PACSTALL_DATABASE_PORT", 3306)
	Database.User = env.GetEnvStringOrDefault("PACSTALL_DATABASE_USER", "root")
	Database.Password = env.GetEnvStringOrDefault("PACSTALL_DATABASE_PASSWORD", "changeme")
	Database.Name = env.GetEnvStringOrDefault("PACSTALL_DATABASE_NAME", "pacstall")
}

// Configuration for the Matomo API
var Matomo = struct {
	Enabled bool
}{}

func initMatomoEnv() {
	Matomo.Enabled = env.GetEnvBoolOrDefault("PACSTALL_MATOMO_ENABLED", false)
}

// Configuration for the Repology API
var Repology = struct {
	Enabled bool
}{}

func initRepologyEnv() {
	Repology.Enabled = env.GetEnvBoolOrDefault("PACSTALL_REPOLOGY_ENABLED", false)
}

func Init() {
	initPacNexusEnv()
	initPacSightEnv()
	initDatabaseEnv()
	initDiscordEnv()
	initMatomoEnv()
	initPacstallProgramsEnv()
	initRepologyEnv()
}
