VERSION="development"

server/dist: $(shell find server -not \( -path server/tmp -prune \) -not \( -path server/dist -prune \) -type f)
	which go
	+$(MAKE) -s -C server

client/dist: $(shell find client -not \( -path client/dist -prune \) -not \( -path client/.parcel-cache -prune \) -type f)
	which node
	+$(MAKE) -s -C client

dist:
	mkdir -p dist
	mkdir -p dist/public
	[ -d ./dist/pacstall-programs ] && rm -rf dist/pacstall-programs || :
	cp -r client/dist/* dist/public
	cp -r server/dist/* dist

docker:
	docker build --rm --build-arg VITE_VERSION="${VERSION}" --no-cache -t webserver .
	docker tag webserver "ghcr.io/pacstall/webserver:${VERSION}"
	docker tag webserver ghcr.io/pacstall/webserver:latest


#### Commands
.PHONY: run clean version

run: server/dist client/dist dist
	cd dist && ./webserver

clean:
	cd server && make clean
	if [ -d dist ]; then rm -rf dist; fi
	if [ -d client/dist ]; then rm -rf client/dist; fi
	if [ -d client/.parcel-cache ]; then rm -rf client/.parcel-cache; fi

prepare: .git/hooks/pre-commit

.git/hooks/pre-commit: ./hooks/client.sh
	cp ./hooks/client.sh .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
