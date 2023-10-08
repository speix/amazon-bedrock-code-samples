# Load all environment variables from the .env file.
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

BASE_DIR=$(shell pwd)

# Install go dependencies.
dep:
	go get -v -d ./...

# Build binaries to be run locally.
build-bedrock: dep
	go build -v -o bin/bedrock src/cmd/bedrock/main.go

run-bedrock: build-bedrock
	./bin/bedrock --debug

# Ensure a command exists.
cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
		(echo "ERROR: '$(*)' must be installed and available on your PATH."; exit 1)