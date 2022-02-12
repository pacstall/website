all: webpacd.tar.gz

### Binaries

webpacd.tar.gz: redist
	[ -d ./redist/pacstall-programs ] && rm -rf redist/pacstall-programs || :
	cd redist && tar -zcvf release *
	mv ./redist/release ./webpacd.tar.gz 

server/bin:
	which go
	cd server && make

client/dist:
	which node
	npm i -g yarn
	cd client && yarn && yarn clean && yarn build

redist: server/bin client/dist
	mkdir -p redist
	mkdir -p redist/public
	[ -d ./redist/pacstall-programs ] && rm -rf redist/pacstall-programs || :
	git clone https://github.com/pacstall/pacstall-programs redist/pacstall-programs
	cp -r client/dist/* redist/public
	cp -r server/bin/* redist


#### Commands
.PHONY: run clean

run: redist run-dist
	cd redist && ./webpacd

clean:
	cd server && make clean
	if [ -f deps_ok ]; then rm deps_ok; fi
	if [ -d redist ]; then rm -rf redist; fi
	if [ -d client/dist ]; then rm -rf client/dist; fi
	if [ -d client/.parcel-cache ]; then rm -rf client/.parcel-cache; fi
	if [ -d client/node_modules ]; then rm -rf client/node_modules; fi
	if [ -f webpacd.tar.gz ]; then rm webpacd.tar.gz; fi