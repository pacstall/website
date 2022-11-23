package config

import (
	"os"
	"strings"
)

var Discord = struct {
	Token string
	ChannelID string
	Enabled bool
	Tags string
} {
	Token: os.Getenv("PACSTALL_DISCORD_TOKEN"),
	ChannelID: os.Getenv("PACSTALL_DISCORD_CHANNEL_ID"),
	Enabled: func (isEnabled string) bool {
		return isEnabled == "1" || isEnabled == "true"
	}(strings.TrimSpace(os.Getenv("PACSTALL_DISCORD_ENABLED"))),
	Tags: os.Getenv("PACSTALL_DISCORD_TAGS"),
}