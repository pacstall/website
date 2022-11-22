package config

import (
	"pacstall.dev/webserver/types/config"
)

var Logging = config.LoggingConfig{}

func setLogging(conf tomlLoggingConfig) {
	Logging.DiscordToken = conf.DiscordToken
	Logging.DiscordChannelID = conf.DiscordChannelID
	Logging.DiscordEnabled = conf.DiscordEnabled
	Logging.DiscordTags = conf.DiscordTags
}
