all: clean gen bin/pacserver

run:
	go run .
bin/pacserver: $(shell find ./ -name '*.go')
	go build -ldflags "-w -s" .
	mkdir -p bin
	mv ./website ./bin/webpacd
	cp ./webpacd.toml.dist ./bin
clean:
	[ -d ./bin ] && rm -r bin || :

gen:
	go run ./codegen/main.go
