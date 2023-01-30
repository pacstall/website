all: dist

dist: $(shell find . -not \( -path dist -prune \) -not \( -path .parcel-cache -prune \) -type f)
	npm install
	npm run lint
	npm run clean
	npm run build