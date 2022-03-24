VERSION = 2.0.0

NPROCS = $(shell grep -c 'processor' /proc/cpuinfo)
MAKEFLAGS += -j$(NPROCS)

all: webpacd.tar.gz

### Binaries

webpacd.tar.gz: redist
	[ -d ./redist/pacstall-programs ] && rm -rf redist/pacstall-programs || :
	cd redist && tar -zcvf release *
	mv ./redist/release ./webpacd.tar.gz 

server/bin:
	which go
	$(MAKE) -s -C server

client/dist:
	which node
	$(MAKE) -s -C client	

redist: server/bin client/dist
	mkdir -p redist
	mkdir -p redist/public
	[ -d ./redist/pacstall-programs ] && rm -rf redist/pacstall-programs || :
	git clone https://github.com/pacstall/pacstall-programs redist/pacstall-programs
	cp -r client/dist/* redist/public
	cp -r server/bin/* redist


#### Commands
.PHONY: run clean version

run: redist
	cd redist && ./webpacd

clean:
	cd server && make clean
	if [ -d redist ]; then rm -rf redist; fi
	if [ -d client/dist ]; then rm -rf client/dist; fi
	if [ -d client/.parcel-cache ]; then rm -rf client/.parcel-cache; fi
	if [ -f webpacd.tar.gz ]; then rm webpacd.tar.gz; fi

version:
	@echo "$(VERSION)"
