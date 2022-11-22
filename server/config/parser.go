package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"pacstall.dev/webserver/log"

	"github.com/BurntSushi/toml"
)

const defaultConfigPath = "./webserver.toml"

var configPath = flag.String("config", defaultConfigPath, fmt.Sprintf("Path to configuration file. Default: %s", defaultConfigPath))
var IsProduction = false

func Load() {
	flag.Parse()
	cfg := loadConfig()
	setLogging(cfg.Logging)
	setTCPServer(cfg.TCPServer)
	setPacstallPrograms(cfg.PacstallPrograms)
	IsProduction = cfg.Production

	log.Info("Loaded configuration from %v", *configPath)
	log.Info("Running in %v mode", func() string {
		mode := "development"
		if IsProduction {
			mode = "production"
		}
		return mode
	}())
	log.Debug("Logging configuration: %#v", Logging)
	log.Debug("Server configuration: %#v", TCPServer)
	log.Debug("Pacstall Programs configuration: %#v", PacstallPrograms)
}

func loadConfig() tomlConfiguration {
	data := tomlConfiguration{}
	bytes, err := os.ReadFile(*configPath)
	if err != nil {
		log.Error("Could not read file '%s'\n%v", *configPath, err)
	}

	if err = toml.Unmarshal(bytes, &data); err != nil {
		log.Error("Could not parse file '%s'\n%v", *configPath, err)
	}

	validate(data)
	return data
}

func prettify(data tomlConfiguration) string {
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Warn(err.Error())
	}
	return string(out)
}

func validate(data tomlConfiguration) {
	config_error := false

	defer func() {
		if config_error {
			os.Exit(1)
		}
	}()

	if data.PacstallPrograms.URL == "" {
		log.Error("Configuration file '%s' is missing required attribute `pacstall_programs.path\n`", *configPath)
		config_error = true

	}

	if data.PacstallPrograms.TempDir == "" {
		log.Error("Configuration file '%s' is missing required attribute `pacstall_programs.tmp_dir\n`", *configPath)
		config_error = true

	}

	if data.PacstallPrograms.UpdateInterval == 0 {
		log.Error("Configuration file '%s' is missing required attribute `pacstall_programs.update_interval\n`", *configPath)
		config_error = true

	}

	if data.PacstallPrograms.MaxOpenFiles == 0 {
		log.Error("Configuration file '%s' is missing required attribute `pacstall_programs.max_open_files\n`", *configPath)
		config_error = true

	}

	if data.TCPServer.Port == 0 {
		log.Error("Configuration file '%s' is missing required attribute `tcp_server.port\n`", *configPath)
		config_error = true

	}

	if data.TCPServer.PublicDir == "" {
		log.Error("Configuration file '%s' is missing required attribute `tcp_server.public_dir\n`", *configPath)
		config_error = true

	}
}
