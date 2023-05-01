package config

import (
	"os"
	"strings"
)

var Discord = struct {
	Token     string
	ChannelID string
	Enabled   bool
	Tags      string
}{
	Token:     os.Getenv("PACSTALL_DISCORD_TOKEN"),
	ChannelID: os.Getenv("PACSTALL_DISCORD_CHANNEL_ID"),
	Enabled: func(isEnabled string) bool {
		return isEnabled == "1" || isEnabled == "true"
	}(strings.TrimSpace(os.Getenv("PACSTALL_DISCORD_ENABLED"))),
	Tags: os.Getenv("PACSTALL_DISCORD_TAGS"),
}

var Database = struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}{
	Host: os.Getenv("PACSTALL_DATABASE_HOST"),
	Port: func(port string) int {
		if port == "" {
			return 3306
		}

		return toInt(port)
	}(os.Getenv("PACSTALL_DATABASE_PORT")),
	User:     os.Getenv("PACSTALL_DATABASE_USER"),
	Password: os.Getenv("PACSTALL_DATABASE_PASSWORD"),
	Name:     os.Getenv("PACSTALL_DATABASE_NAME"),
}
