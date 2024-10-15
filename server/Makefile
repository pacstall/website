LDFLAGS= \
	-w \
	-s \
	-X 'pacstall.dev/webserver/pkg/common/config/build.Production=true' \
	-X 'pacstall.dev/webserver/pkg/common/config/build.Version=${VERSION}'

all: dist/pacnexus dist/pacsight

test:
	GO_ENV=test go test -v types/pac/parser/parse_test.go

test_internal:
	GO_ENV=test go test -v types/pac/parser/pacsh/internal/git_version_test.go

run-database:
	@(cd .. && docker compose up -d mariadb)

run-pacnexus: run-database
	@go run cmd/pacnexus/main.go

run-pacsight:
	@go run cmd/pacsight/main.go


run:
	@trap 'kill 0' SIGINT; \
	$(MAKE) run-pacsight | sed 's/^/[pacsight] /' & \
	$(MAKE) run-pacnexus | sed 's/^/[pacnexus] /' & \
	wait
	@echo "Both executables completed or interrupted."

dist/pacnexus: $(shell find . -not \( -path ./tmp -prune \) -not \( -path ./dist -prune \) -type f)
	CGO_ENABLED=0 go build -o dist/pacnexus -ldflags "${LDFLAGS}" cmd/pacnexus/main.go

dist/pacsight: $(shell find . -not \( -path ./tmp -prune \) -not \( -path ./dist -prune \) -type f)
	CGO_ENABLED=0 go build -o dist/pacsight -ldflags "${LDFLAGS}" cmd/pacsight/main.go

clean:
	@[ -d ./dist ] && rm -rf dist || :
	@[ -d ./programs ] && rm -rf programs || :
	@[ -d ./tmp ] && rm -rf tmp || :
	@echo "Cleaned up."

fmt:
 