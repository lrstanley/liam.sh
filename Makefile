.DEFAULT_GOAL := build-all

export PROJECT := "httpserver"
export PACKAGE := "github.com/lrstanley/liam.sh/cmd/httpserver"
export COMPOSE_PROJECT := "liamsh"
export COMPOSE_ARGS := "postgres"
export COMPOSE_DOCKER_CLI_BUILD := 1
export DOCKER_BUILDKIT := 1
export USER := $(shell id -u)
export GROUP := $(shell id -g)
export LICENSE_IGNORE := "graphql-tag"

license:
	curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

build-all: clean node-fetch go-fetch node-build go-build
	@echo

up: node-upgrade-deps go-upgrade-deps
	@echo

clean:
	/bin/rm -rfv "cmd/httpserver/public/dist/*" ${PROJECT}

docker:
	docker compose \
		--project-name ${COMPOSE_PROJECT} \
		--file docker-compose.yaml \
		up \
		--remove-orphans \
		--build \
		--timeout 0 ${COMPOSE_ARGS}

docker-clean:
	docker compose \
		--project-name ${COMPOSE_PROJECT} \
		--file docker-compose.yaml \
		down \
		--volumes \
		--remove-orphans \
		--rmi local --timeout 1

docker-build:
	docker build \
		--tag ${PROJECT} \
		--force-rm .

# frontend
node-fetch:
	command -v pnpm >/dev/null >&2 || npm install \
		--no-audit \
		--no-fund \
		--quiet \
		--global pnpm
	cd cmd/httpserver/public && \
		pnpm install --silent

node-upgrade-deps:
	cd cmd/httpserver/public && \
		pnpm up -i

node-prepare: node-fetch
	cd cmd/httpserver/public && \
		pnpm exec graphql-codegen --config graphql.yaml

node-lint: node-build # needed to generate eslint auto-import ignores.
	cd cmd/httpserver/public && \
		pnpm exec eslint \
			--ignore-path ../../../.gitignore \
			--ext .js,.ts,.vue .
	cd cmd/httpserver/public && \
		pnpm exec vue-tsc --noEmit

node-debug: node-prepare
	cd cmd/httpserver/public && \
		pnpm exec vite

node-build: node-prepare
	cd cmd/httpserver/public && \
		pnpm exec vite build

node-preview: node-build
	cd cmd/httpserver/public && \
		pnpm exec vite preview

# backend
go-prepare:
	go generate -x ./...

go-fetch:
	go mod download
	go mod tidy

go-upgrade-deps:
	go get -u ./...
	go mod tidy

go-upgrade-deps-patch:
	go get -u=patch ./...
	go mod tidy

go-dlv: go-prepare
	dlv debug \
		--headless --listen=:2345 \
		--api-version=2 --log \
		--allow-non-terminal-interactive \
		${PACKAGE} -- --debug

go-debug: go-prepare
	go run ${PACKAGE} --debug

go-build: go-prepare go-fetch
	CGO_ENABLED=0 \
	go build \
		-ldflags '-d -s -w -extldflags=-static' \
		-tags=netgo,osusergo,static_build \
		-installsuffix netgo \
		-buildvcs=false \
		-trimpath \
		-o ${PROJECT} \
		${PACKAGE}
