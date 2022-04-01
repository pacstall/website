package config

type TCPServerConfig struct {
	Port      uint16
	PublicDir string
}

var TCPServer = TCPServerConfig{}

func setTCPServer(conf tomlTCPServerConfig) {
	TCPServer.Port = conf.Port
	TCPServer.PublicDir = conf.PublicDir
}
