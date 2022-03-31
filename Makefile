.DEFAULT_GOAL := build

BINARY=liam.sh

$(info $(shell mkdir -p $(DIRS)))

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

fetch: ## Fetches the necessary dependencies to build.
	which rice 2>&1 > /dev/null || go install github.com/GeertJohan/go.rice/rice@latest
	go mod download
	go mod tidy

upgrade-deps: ## Upgrade all dependencies to the latest version.
	go get -u ./...

upgrade-deps-patch: ## Upgrade all dependencies to the latest patch release.
	go get -u=patch ./...

clean: ## Cleans up generated files/folders from the build.
	/bin/rm -rfv "dist/" "${BINARY}" rice-box.go

build: fetch clean ## Compile and generate a binary with static assets embedded.
	rice -v embed-go
	go build -ldflags '-d -s -w' -tags netgo -installsuffix netgo -v -o "${BINARY}"

debug: fetch clean
	go run *.go run --debug --git.release-skip --chat-link https://some-url.com/chat
