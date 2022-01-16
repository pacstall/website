package cfg

import (
	"encoding/json"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

const CONFIG_PATH = "./config.toml"

type configuration struct {
	TCPServer        tcpServerConfig        `toml:"tcp_server"`
	PacstallPrograms pacstallProgramsConfig `toml:"pacstall_programs"`
}

type tcpServerConfig struct {
	Port int `toml:"port"`
}

type pacstallProgramsConfig struct {
	Path           string `toml:"path"`
	TempDir        string `toml:"tmp_dir"`
	UpdateInterval int    `toml:"update_interval"`
}

var Config configuration = loadConfig()

func loadConfig() configuration {
	data := configuration{}
	bytes, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		log.Panicf("Could not read file 'config.toml'\n%v", err)
	}

	if err = toml.Unmarshal(bytes, &data); err != nil {
		log.Panicf("Could not parse file 'config.toml'\n%v", err)
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
	if data.PacstallPrograms.Path == "" {
		log.Fatalln("Configuration file 'config.toml' is missing required attribute `pacstall_programs.path`")
	}

	if data.PacstallPrograms.TempDir == "" {
		log.Fatalln("Configuration file 'config.toml' is missing required attribute `pacstall_programs.tmp_dir`")
	}

	if data.PacstallPrograms.UpdateInterval == 0 {
		log.Fatalln("Configuration file 'config.toml' is missing required attribute `pacstall_programs.update_interval`")
	}

	if data.TCPServer.Port == 0 {
		log.Fatalln("Configuration file 'config.toml' is missing required attribute `tcp_server.port`")
	}
}
