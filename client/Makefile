all: dist dist/swagger

dist: $(shell find . -not \( -path dist -prune \) -not \( -path .parcel-cache -prune \) -type f)
	npm ci
	npm run lint
	npm run clean
	npm run build

dist/swagger: dist
	mv dist/public/public/swagger dist/swagger