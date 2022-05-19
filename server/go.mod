module pacstall.dev/webserver

go 1.18

require (
	github.com/BurntSushi/toml v1.0.0 // for parsing toml config files
	github.com/fatih/color v1.13.0 // for colorizing output
	github.com/gorilla/mux v1.8.0 // for http request routing
	github.com/hashicorp/go-version v1.5.0 // for version parsing
)

require github.com/schollz/progressbar/v3 v3.8.6

require (
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
)

require (
	github.com/k0kubun/go-ansi v0.0.0-20180517002512-3bf9e2903213
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/sys v0.0.0-20220403205710-6acee93ad0eb // indirect
)
