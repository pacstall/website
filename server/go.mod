module pacstall.dev/webserver

go 1.18

require (
	github.com/BurntSushi/toml v1.2.0 // for parsing toml config files
	github.com/fatih/color v1.13.0 // for colorizing output
	github.com/gorilla/mux v1.8.0 // for http request routing
	github.com/hashicorp/go-version v1.6.0 // for version parsing
)

require github.com/schollz/progressbar/v3 v3.11.0

require golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect

require (
	github.com/k0kubun/go-ansi v0.0.0-20180517002512-3bf9e2903213
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
)
