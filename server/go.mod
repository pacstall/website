module pacstall.dev/webserver

go 1.18

require (
	github.com/BurntSushi/toml v1.2.0 // for parsing toml config files
	github.com/fatih/color v1.13.0 // for colorizing output
	github.com/gorilla/mux v1.8.0 // for http request routing
	github.com/hashicorp/go-version v1.6.0 // for version parsing
)

require github.com/schollz/progressbar/v3 v3.12.0

require golang.org/x/term v0.1.0 // indirect

require (
	github.com/k0kubun/go-ansi v0.0.0-20180517002512-3bf9e2903213
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	golang.org/x/sys v0.1.0 // indirect
)
