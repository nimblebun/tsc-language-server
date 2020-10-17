default: build

build:
	mkdir -p dist
	python scripts/generateDefaultConfig.py
	go build -o dist

.PHONY: build
