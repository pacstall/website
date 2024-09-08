all: dist

dist: $(shell find . -not \( -path dist -prune \) -not \( -path .parcel-cache -prune \) -type f)
	npm install
	npm run lint
	npm run clean
	VITE_VERSION=$$VITE_VERSION npm run build

run:
	npm install
	npm run lint:fix
	npm start
