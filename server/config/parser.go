package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

const CONFIG_PATH = "./webserver.toml"

type configuration struct {
	TCPServer        tcpServerConfig        `toml:"tcp_server"`
	PacstallPrograms pacstallProgramsConfig `toml:"pacstall_programs"`
	Production       bool                   `toml:"production"`
	FeatureFlags     featureFlagsConfig     `toml:"feature_flags"`
}

type featureFlagsConfig struct {
	OldSyntax          bool                           `toml:"old_syntax"`
	PackageDetailsPage packageDetailsPageFeatureFlags `toml:"package_details_page"`
}

type packageDetailsPageFeatureFlags struct {
	LastUpdated     bool `toml:"last_updated"`
	Votes           bool `toml:"votes"`
	Popularity      bool `toml:"popularity"`
	InstallProtocol bool `toml:"install_protocol"`
	Comments        bool `toml:"comments"`
}

type tcpServerConfig struct {
	Port      int    `toml:"port"`
	PublicDir string `toml:"public_dir"`
}

type pacstallProgramsConfig struct {
	Path           string `toml:"path"`
	TempDir        string `toml:"tmp_dir"`
	UpdateInterval int    `toml:"update_interval"`
	MaxOpenFiles   int    `toml:"max_open_files"`
}

var Config configuration = loadConfig()

func loadConfig() configuration {
	data := configuration{}
	bytes, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		log.Fatalf("Could not read file '%s'\n%v", CONFIG_PATH, err)
	}

	if err = toml.Unmarshal(bytes, &data); err != nil {
		log.Fatalf("Could not parse file '%s'\n%v", CONFIG_PATH, err)
	}

	validate(data)

	log.Printf("Configuration successfully loaded %v", prettify(data))

	return data
}

func prettify(data configuration) string {
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(out)
}

func validate(data configuration) {
	config_error := false

	defer func() {
		if config_error {
			os.Exit(1)
		}
	}()

	if data.PacstallPrograms.Path == "" {
		log.Println("Configuration file 'config.toml' is missing required attribute `pacstall_programs.path`")
	}

	if data.PacstallPrograms.TempDir == "" {
		log.Println("Configuration file 'config.toml' is missing required attribute `pacstall_programs.tmp_dir`")
	}

	if data.PacstallPrograms.UpdateInterval == 0 {
		log.Println("Configuration file 'config.toml' is missing required attribute `pacstall_programs.update_interval`")
	}

	if data.PacstallPrograms.MaxOpenFiles == 0 {
		log.Println("Configuration file 'config.toml' is missing required attribute `pacstall_programs.max_open_files`")
	}

	if data.TCPServer.Port == 0 {
		log.Println("Configuration file 'config.toml' is missing required attribute `tcp_server.port`")
	}

	if data.TCPServer.PublicDir == "" {
		log.Println("Configuration file 'config.toml' is missing required attribute `tcp_server.public_dir`")
	}
}
