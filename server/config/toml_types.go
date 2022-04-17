package config

type tomlConfiguration struct {
	TCPServer        tomlTCPServerConfig        `toml:"tcp_server"`
	PacstallPrograms tomlPacstallProgramsConfig `toml:"pacstall_programs"`
	Production       bool                       `toml:"production"`
	FeatureFlags     tomlFeatureFlagsConfig     `toml:"feature_flags"`
	Logging          tomlLoggingConfig          `toml:"logging"`
}

type tomlFeatureFlagsConfig struct {
	OldSyntax          bool                               `toml:"old_syntax"`
	PackageDetailsPage tomlPackageDetailsPageFeatureFlags `toml:"package_details_page"`
}

type tomlLoggingConfig struct {
	FancyLogs bool   `toml:"fancy_logs"`
	LogLevel  string `toml:"log_level"`
}

type tomlPackageDetailsPageFeatureFlags struct {
	LastUpdated     bool `toml:"last_updated"`
	Votes           bool `toml:"votes"`
	Popularity      bool `toml:"popularity"`
	InstallProtocol bool `toml:"install_protocol"`
	Comments        bool `toml:"comments"`
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
