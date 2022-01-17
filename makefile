all: dependency-check clean build-server build-client build-redist

run: all run-dist

dev-client:
	@cd client && yarn && yarn start
dev-server:
	@cd server && make run

build-server:
	@cd server && make

build-client:
	@cd client && yarn && yarn clean && yarn build

build-redist:
	@mkdir -p redist
	@mkdir -p redist/public
	@[ -d ./redist/pacstall-programs ] && rm -rf redist/pacstall-programs || :
	@git clone https://github.com/pacstall/pacstall-programs redist/pacstall-programs
	@cp -r client/dist/* redist/public
	@cp -r server/bin/* redist

clean:
	@cd server && make clean
run-dist:
	@cd redist && ./webpacd

dependency-check:
	@echo "if build dependency 'node' is missing, make will error here"; which node >/dev/null
	@echo "dependency 'node': found\n"
	@npm i -g yarn; echo "dependency 'yarn': found\n"
	@echo "if build dependency 'go' is missing, make will error here"; which go >/dev/null
	@echo "dependency 'go': found\n"