.DEFAULT_GOAL := build

DIRS=bin dist
BINARY=liam.sh

$(info $(shell mkdir -p $(DIRS)))
BIN=$(CURDIR)/bin
export GOBIN=$(CURDIR)/bin

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

docker-build: fetch clean ## Compile within a docker container (no go or other dependencies required)
	docker build --rm --force-rm -f Dockerfile -t lrstanley/liam.sh:latest .

docker-push: ## Push docker images to <version> and latest
	docker rmi $(docker images -f "dangling=true" -q) | true
	docker push lrstanley/liam.sh:latest

fetch: ## Fetches the necessary dependencies to build.
	which $(BIN)/rice 2>&1 > /dev/null || go get github.com/GeertJohan/go.rice/rice
	go mod download
	go mod tidy
	go mod vendor

upgrade-deps: ## Upgrade all dependencies to the latest version.
	go get -u ./...

upgrade-deps-patch: ## Upgrade all dependencies to the latest patch release.
	go get -u=patch ./...

clean: ## Cleans up generated files/folders from the build.
	/bin/rm -rfv "dist/" "${BINARY}" rice-box.go

build: fetch clean ## Compile and generate a binary with static assets embedded.
	$(BIN)/rice -v embed-go
	go build -ldflags '-d -s -w' -tags netgo -installsuffix netgo -v -o "${BINARY}"

debug: fetch clean
	go run *.go run --debug --git.release-skip --chat-link https://some-url.com/chat
