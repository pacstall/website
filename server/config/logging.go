package config

import "pacstall.dev/webserver/log"

type LoggingConfig struct {
	FancyLogs bool
	Level     log.Level
}

var Logging = LoggingConfig{}

func setLogging(conf tomlLoggingConfig) {
	Logging.FancyLogs = conf.FancyLogs
	Logging.Level = log.NewLogLevel(conf.LogLevel)
}
