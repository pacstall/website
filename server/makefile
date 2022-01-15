all: clean bin/pacserver

run:
	go run .
bin/pacserver: $(shell find ./ -name '*.go')
	go build .
	mkdir -p bin
	mv ./website ./bin/pacserver
	cp ./config.toml ./bin
clean:
	rm -r bin
