all: dist/webserver

run: 
	go run bin/webserver/main.go
dist/webserver: $(shell find . -not \( -path ./tmp -prune \) -not \( -path ./dist -prune \) -type f)
	CGO_ENABLED=0 go build -o dist/webserver -ldflags "-w -s" bin/webserver/main.go
	cp ./webserver.toml.dist ./dist/webserver.toml
clean:
	[ -d ./dist ] && rm -r dist || :

test:
	go test -v ./...
