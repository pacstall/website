LDFLAGS= \
	-w \
	-s \
	-X 'pacstall.dev/webserver/config/build.UpdateInterval=900' \
	-X 'pacstall.dev/webserver/config/build.TempDir=./tmp' \
	-X 'pacstall.dev/webserver/config/build.MaxOpenFiles=100' \
	-X 'pacstall.dev/webserver/config/build.Port=3300' \
	-X 'pacstall.dev/webserver/config/build.PublicDir=../client/dist' \
	-X 'pacstall.dev/webserver/config/build.Production=true' \
	-X 'pacstall.dev/webserver/config/build.GitURL=https://github.com/pacstall/pacstall-programs.git' \
	-X 'pacstall.dev/webserver/config/build.GitClonePath=./programs' \
	-X 'pacstall.dev/webserver/config/build.Version=${VERSION}'

all: dist/webserver

run: 
	go run bin/webserver/main.go
dist/webserver: $(shell find . -not \( -path ./tmp -prune \) -not \( -path ./dist -prune \) -type f)
	CGO_ENABLED=0 go build -o dist/webserver -ldflags "${LDFLAGS}" bin/webserver/main.go
clean:
	[ -d ./dist ] && rm -r dist || :
test:
	go test -v ./...
fmt:
	go fmt ./...
