LDFLAGS= \
	-w \
	-s \
	-X 'pacstall.dev/webserver/config/build.UpdateInterval=900' \
	-X 'pacstall.dev/webserver/config/build.TempDir=./tmp' \
	-X 'pacstall.dev/webserver/config/build.MaxOpenFiles=100' \
	-X 'pacstall.dev/webserver/config/build.Port=3300' \
	-X 'pacstall.dev/webserver/config/build.PublicDir=./public' \
	-X 'pacstall.dev/webserver/config/build.Production=true' \
	-X 'pacstall.dev/webserver/config/build.GitURL=https://github.com/pacstall/pacstall-programs.git' \
	-X 'pacstall.dev/webserver/config/build.GitClonePath=./programs' \
	-X 'pacstall.dev/webserver/config/build.Version=${VERSION}'

all: dist/webserver

test:
	PACSTALL_DATABASE_HOST=localhost \
	PACSTALL_DATABASE_PORT=3306 \
	PACSTALL_DATABASE_USER=root \
	PACSTALL_DATABASE_PASSWORD=changeme \
	PACSTALL_DATABASE_NAME=pacstall \
	PACSTALL_DISCORD_ENABLED=false \
	PACSTALL_DISCORD_TOKEN="" \
	PACSTALL_DISCORD_CHANNEL_ID="" \
	PACSTALL_DISCORD_TAGS="" \
	PACSTALL_MATOMO_ENABLED="false" \
	PACSTALL_REPOLOGY_ENABLED="false" \
	PACSTALL_PROGRAMS_GIT_BRANCH="master" \
	GO_ENV=test go test -v types/pac/parser/parse_test.go

test_internal:
	PACSTALL_DATABASE_HOST=localhost \
	PACSTALL_DATABASE_PORT=3306 \
	PACSTALL_DATABASE_USER=root \
	PACSTALL_DATABASE_PASSWORD=changeme \
	PACSTALL_DATABASE_NAME=pacstall \
	PACSTALL_DISCORD_ENABLED=false \
	PACSTALL_DISCORD_TOKEN="" \
	PACSTALL_DISCORD_CHANNEL_ID="" \
	PACSTALL_DISCORD_TAGS="" \
	PACSTALL_MATOMO_ENABLED="false" \
	PACSTALL_REPOLOGY_ENABLED="false" \
	PACSTALL_PROGRAMS_GIT_BRANCH="master" \
	GO_ENV=test go test -v types/pac/parser/pacsh/internal/git_version_test.go

run: 
	(cd .. && docker compose up -d mariadb)

	PACSTALL_DATABASE_HOST=localhost \
	PACSTALL_DATABASE_PORT=3306 \
	PACSTALL_DATABASE_USER=root \
	PACSTALL_DATABASE_PASSWORD=changeme \
	PACSTALL_DATABASE_NAME=pacstall \
	PACSTALL_DISCORD_ENABLED=false \
	PACSTALL_DISCORD_TOKEN="" \
	PACSTALL_DISCORD_CHANNEL_ID="" \
	PACSTALL_DISCORD_TAGS="" \
	PACSTALL_MATOMO_ENABLED="false" \
	PACSTALL_REPOLOGY_ENABLED="false" \
	PACSTALL_PROGRAMS_GIT_BRANCH="5.0.0-master" \
	go run bin/webserver/main.go

dist/webserver: $(shell find . -not \( -path ./tmp -prune \) -not \( -path ./dist -prune \) -type f)
	CGO_ENABLED=0 go build -o dist/webserver -ldflags "${LDFLAGS}" bin/webserver/main.go
clean:
	[ -d ./dist ] && rm -r dist || :

fmt:
	go fmt ./...
