VERSION = 2.1.0-rc5

NPROCS = $(shell grep -c 'processor' /proc/cpuinfo)
MAKEFLAGS += -j$(NPROCS)

server/dist:
	which go
	+$(MAKE) -s -C server

client/dist:
	which node
	+$(MAKE) -s -C client	

dist:
	mkdir -p dist
	mkdir -p dist/public
	[ -d ./dist/pacstall-programs ] && rm -rf dist/pacstall-programs || :
	cp -r client/dist/* dist/public
	cp -r server/dist/* dist
#git clone https://github.com/pacstall/pacstall-programs dist/pacstall-programs


#### Commands
.PHONY: run clean version

run: dist
	cd dist && ./webserver

clean:
	cd server && make clean
	if [ -d dist ]; then rm -rf dist; fi
	if [ -d client/dist ]; then rm -rf client/dist; fi
	if [ -d client/.parcel-cache ]; then rm -rf client/.parcel-cache; fi
	if [ -f webserver.tar.gz ]; then rm webserver.tar.gz; fi

version:
	@echo "$(VERSION)"
