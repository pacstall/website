all: clean build-server build-client build-redist

dev-client:
	cd client && yarn && yarn dev
dev-server:
	cd server && make run

build-server:
	cd server && make

build-client:
	NODE_ENV="production" cd client && npm install && npm run build

build-redist:
	mkdir -p redist
	mkdir -p redist/public
	[ -d ./redist/pacstall-programs ] && rm -rf redist/pacstall-programs || :
	git clone https://github.com/pacstall/pacstall-programs redist/pacstall-programs
	cp -r client/out/* redist/public
	cp -r server/bin/* redist

clean:
	cd server && make clean
