package config

import (
	"time"

	"pacstall.dev/webserver/config/build"
	"pacstall.dev/webserver/config/internal"
)

type GlobalConfiguration struct {
	ServerConfiguration           ServerConfiguration
	DatabaseConfiguration         DatabaseConfiguration
	DiscordConfiguration          DiscordConfiguration
	PacstallProgramsConfiguration PacstallProgramsConfiguration
	MatomoConfiguration           MatomoConfiguration
	RepologyConfiguration         RepologyConfiguration
}

func Parse() GlobalConfiguration {
	server := ServerConfiguration{
		Port:         internal.GetEnvIntOrDefault("PACSTALL_SERVER_PORT", 3300),
		PublicDir:    internal.GetEnvStringOrDefault("PACSTALL_SERVER_PUBLIC_DIR", "../client/dist"),
		Version:      build.Version,
		TempDir:      internal.GetEnvStringOrDefault("PACSTALL_SERVER_TEMP_DIR", "./tmp"),
		MaxOpenFiles: internal.GetEnvIntOrDefault("PACSTALL_SERVER_MAX_OPEN_FILES", 100),
		Production:   internal.GetEnvBoolOrDefault("PACSTALL_SERVER_PRODUCTION", false),
	}

	discord := parseDiscordConfiguration()

	programs := PacstallProgramsConfiguration{
		Branch:         internal.GetEnvStringOrDefault("PACSTALL_PROGRAMS_GIT_BRANCH", "master"),
		RepositoryUrl:  internal.GetEnvStringOrDefault("PACSTALL_PROGRAMS_GIT_REPOSITORY", "https://github.com/pacstall/pacstall-programs.git"),
		UpdateInterval: time.Duration(internal.GetEnvIntOrDefault("PACSTALL_PROGRAMS_UPDATE_INTERVAL", 900)) * time.Second,
		ClonePath:      internal.GetEnvStringOrDefault("PACSTALL_PROGRAMS_CLONE_PATH", "./programs"),
	}

	database := DatabaseConfiguration{
		Host:     internal.GetEnvString("PACSTALL_DATABASE_HOST"),
		Port:     internal.GetEnvInt("PACSTALL_DATABASE_PORT"),
		User:     internal.GetEnvString("PACSTALL_DATABASE_USER"),
		Password: internal.GetEnvString("PACSTALL_DATABASE_PASSWORD"),
		Name:     internal.GetEnvString("PACSTALL_DATABASE_NAME"),
	}

	matomo := MatomoConfiguration{
		Enabled: internal.GetEnvBoolOrDefault("PACSTALL_MATOMO_ENABLED", false),
	}

	repology := RepologyConfiguration{
		Enabled:        internal.GetEnvBoolOrDefault("PACSTALL_REPOLOGY_ENABLED", true),
		UpdateInterval: time.Duration(internal.GetEnvIntOrDefault("PACSTALL_REPOLOGY_UPDATE_INTERVAL", 43200)) * time.Second,
	}

	return GlobalConfiguration{
		ServerConfiguration:           server,
		DatabaseConfiguration:         database,
		DiscordConfiguration:          discord,
		PacstallProgramsConfiguration: programs,
		MatomoConfiguration:           matomo,
		RepologyConfiguration:         repology,
	}
}

type ServerConfiguration struct {
	Production   bool
	Port         int
	PublicDir    string
	Version      string
	TempDir      string
	MaxOpenFiles int
}

// Configuration for the discord integration
type DiscordConfiguration struct {
	Token     string
	ChannelID string
	Enabled   bool
	Tags      string
}

func parseDiscordConfiguration() DiscordConfiguration {
	enabled := internal.GetEnvBoolOrDefault("PACSTALL_DISCORD_ENABLED", false)
	if !enabled {
		return DiscordConfiguration{
			Enabled: false,
		}
	}

	return DiscordConfiguration{
		Enabled:   true,
		Token:     internal.GetEnvString("PACSTALL_DISCORD_TOKEN"),
		ChannelID: internal.GetEnvStringOrDefault("PACSTALL_DISCORD_CHANNEL_ID", ""),
		Tags:      internal.GetEnvString("PACSTALL_DISCORD_TAGS"),
	}
}

type PacstallProgramsConfiguration struct {
	Branch         string
	RepositoryUrl  string
	UpdateInterval time.Duration
	ClonePath      string
}

// Configuration for the database
type DatabaseConfiguration struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// Configuration for the Matomo API
type MatomoConfiguration struct {
	Enabled bool
}

// Configuration for the Repology API
type RepologyConfiguration struct {
	Enabled        bool
	UpdateInterval time.Duration
}
