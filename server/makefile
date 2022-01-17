all: clean gen bin/pacserver

run:
	go run .
bin/pacserver: $(shell find ./ -name '*.go')
	go build .
	mkdir -p bin
	mv ./website ./bin/pacserver
	cp ./config.toml ./bin/config.toml.dist
clean:
	[ -d ./bin ] && rm -r bin || :

gen:
	go run ./codegen/main.go
