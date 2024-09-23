package env

import (
	"time"

	"pacstall.dev/webserver/pkg/common/config/build"
)

var PacSight = struct {
	Port int
}{}

func initPacSightEnv() {
	PacSight.Port = GetEnvIntOrDefault("PACSTALL_PACSIGHT_PORT", 8080)
}

var PacNexus = struct {
	Port       int
	PublicDir  string
	Version    string
	Production bool
}{}

func initPacNexusEnv() {
	PacNexus.Port = GetEnvIntOrDefault("PACSTALL_PACNEXUS_PORT", 3300)
	PacNexus.PublicDir = GetEnvStringOrDefault("PACSTALL_PACNEXUS_PUBLIC_DIR", "./public")
	PacNexus.Version = build.Version
	PacNexus.Production = toBool(build.Production)
}

// Configuration for the discord integration
var Discord = struct {
	Token     string
	ChannelID string
	Enabled   bool
	Tags      string
}{}

func initDiscordEnv() {
	Discord.Enabled = GetEnvBoolOrDefault("PACSTALL_DISCORD_ENABLED", false)

	if !Discord.Enabled {
		return
	}

	Discord.Token = GetEnvString("PACSTALL_DISCORD_TOKEN")
	Discord.ChannelID = GetEnvString("PACSTALL_DISCORD_CHANNEL_ID")
	Discord.Tags = GetEnvString("PACSTALL_DISCORD_TAGS")
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
	PacstallPrograms.Branch = GetEnvStringOrDefault("PACSTALL_PROGRAMS_GIT_BRANCH", "master")
	PacstallPrograms.UpdateInterval = time.Duration(GetEnvIntOrDefault("PACSTALL_PROGRAMS_UPDATE_INTERVAL", 15*60)) * time.Second
	PacstallPrograms.TempDir = GetEnvStringOrDefault("PACSTALL_PROGRAMS_TEMP_DIR", "./tmp")
	PacstallPrograms.MaxOpenFiles = GetEnvIntOrDefault("PACSTALL_PROGRAMS_MAX_OPEN_FILES", 100)
	PacstallPrograms.GitURL = GetEnvStringOrDefault("PACSTALL_PROGRAMS_GIT_URL", "https://github.com/pacstall/pacstall-programs.git")
	PacstallPrograms.GitClonePath = GetEnvStringOrDefault("PACSTALL_PROGRAMS_GIT_CLONE_PATH", "./programs")
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
	Database.Host = GetEnvStringOrDefault("PACSTALL_DATABASE_HOST", "localhost")
	Database.Port = GetEnvIntOrDefault("PACSTALL_DATABASE_PORT", 3306)
	Database.User = GetEnvStringOrDefault("PACSTALL_DATABASE_USER", "root")
	Database.Password = GetEnvStringOrDefault("PACSTALL_DATABASE_PASSWORD", "changeme")
	Database.Name = GetEnvStringOrDefault("PACSTALL_DATABASE_NAME", "pacstall")
}

// Configuration for the Matomo API
var Matomo = struct {
	Enabled bool
}{}

func initMatomoEnv() {
	Matomo.Enabled = GetEnvBoolOrDefault("PACSTALL_MATOMO_ENABLED", false)
}

// Configuration for the Repology API
var Repology = struct {
	Enabled                bool
	RepologyUpdateInterval time.Duration
}{}

func initRepologyEnv() {
	Repology.Enabled = GetEnvBoolOrDefault("PACSTALL_REPOLOGY_ENABLED", false)

	if !Repology.Enabled {
		return
	}

	Repology.RepologyUpdateInterval = time.Duration(GetEnvIntOrDefault("PACSTALL_REPOLOGY_UPDATE_INTERVAL", 60*60*6)) * time.Second
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
