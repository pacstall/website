all: clean bin/pacserver

run:
	go run bin/webserver/main.go
bin/pacserver: $(shell find ./ -name '*.go')
	go build -ldflags "-w -s" .
	mkdir -p bin
	mv ./website ./bin/webpacd
	cp ./webpacd.toml.dist ./bin
clean:
	[ -d ./bin ] && rm -r bin || :

test:
	grc go test -v ./...
