all: clean test dist/webserver

run:
	go run bin/webserver/main.go
dist/webserver: $(shell find ./ -name '*.go')
	go build -o dist/webserver -ldflags "-w -s" bin/webserver/main.go
	cp ./webserver.toml.dist ./dist
clean:
	[ -d ./dist ] && rm -r dist || :

test:
	grc go test -v ./...
