package config

var Discord = struct {
	Token     string
	ChannelID string
	Enabled   bool
	Tags      string
}{
	Token:     getEnvString("PACSTALL_DISCORD_TOKEN"),
	ChannelID: getEnvString("PACSTALL_DISCORD_CHANNEL_ID"),
	Enabled:   getEnvBool("PACSTALL_DISCORD_ENABLED"),
	Tags:      getEnvString("PACSTALL_DISCORD_TAGS"),
}

var Database = struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}{
	Host:     getEnvString("PACSTALL_DATABASE_HOST"),
	Port:     getEnvInt("PACSTALL_DATABASE_PORT"),
	User:     getEnvString("PACSTALL_DATABASE_USER"),
	Password: getEnvString("PACSTALL_DATABASE_PASSWORD"),
	Name:     getEnvString("PACSTALL_DATABASE_NAME"),
}

var Matomo = struct {
	Enabled bool
}{
	Enabled: getEnvBool("PACSTALL_MATOMO_ENABLED"),
}

var Repology = struct {
	Enabled bool
}{
	Enabled: getEnvBoolOrDefault("PACSTALL_REPOLOGY_ENABLED", true),
}
