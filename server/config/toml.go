package config

type tomlConfiguration struct {
	TCPServer        tomlTCPServerConfig        `toml:"tcp_server"`
	PacstallPrograms tomlPacstallProgramsConfig `toml:"pacstall_programs"`
	Production       bool                       `toml:"production"`
	Logging          tomlLoggingConfig          `toml:"logging"`
}

type tomlLoggingConfig struct {
	DiscordToken     string `toml:"discord_token"`
	DiscordChannelID string `toml:"discord_channel_id"`
	DiscordEnabled   bool   `toml:"discord_enabled"`
	DiscordTags      string `toml:"discord_tags"`
}

type tomlTCPServerConfig struct {
	Port      uint16 `toml:"port"`
	PublicDir string `toml:"public_dir"`
}

type tomlPacstallProgramsConfig struct {
	URL            string `toml:"url"`
	TempDir        string `toml:"tmp_dir"`
	UpdateInterval int    `toml:"update_interval"`
	MaxOpenFiles   uint8  `toml:"max_open_files"`
}
