.DEFAULT_GOAL := build-all

export PROJECT := "httpserver"
export PACKAGE := "github.com/lrstanley/liam.sh/cmd/httpserver"
export DOCKER_BUILDKIT := 1
export NODE_OPTIONS := "--max-old-space-size=3900"

license:
	# curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

prepare: clean node-prepare node-build go-prepare
	@echo

build-all: prepare go-build
	@echo

up: node-upgrade-deps go-upgrade-deps
	@echo

clean:
	/bin/rm -rfv "cmd/httpserver/public/dist/*" ${PROJECT}

generate-language-colors:
	curl -LsS -o- https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml \
		| yq '{"languages": .} | tojson' \
		| go run -mod=mod github.com/wrouesnel/p2cli/cmd/p2@latest \
			--format=json \
			--template="internal/wakapi/colorconst.go.j2" \
			--output="internal/wakapi/colorconst.go"
	gofmt -l -w internal/wakapi/colorconst.go

docker-build:
	docker build \
		--pull \
		--tag ${PROJECT} \
		--file .github/Dockerfile \
		--force-rm .

# web
node-fetch:
	cd web/ && bun install
	if [ ! -f web/app/components/events/objects/events.ts ]; then \
		bunx openapi-typescript -o web/app/components/events/objects/events.ts https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json; \
	fi

node-upgrade-deps:
	cd web/ && bun update -i

node-prepare: license node-fetch
	@echo

node-lint: node-build # needed to generate eslint auto-import ignores.
	cd web/ && bunx eslint \
		--ignore-path ../../../.gitignore \
		--ext .js,.ts,.vue .
	cd web/ && bunx vue-tsc --noEmit

node-debug: node-prepare
	cd web/ && bun run dev

node-build: node-prepare
	cd web/ && bun run build

node-preview: node-build
	cd web/ && bun run preview

# backend
go-clean:
	rm -rf internal/database/ent .gitapicache/

go-prepare: license
	go generate -x ./...
	go mod tidy

go-upgrade-deps:
	go get -u ./...
	go mod tidy

go-dlv: go-prepare
	dlv debug \
		--headless --listen=:2345 \
		--api-version=2 --log \
		--allow-non-terminal-interactive \
		${PACKAGE} -- --debug

go-debug: go-prepare
	go run ${PACKAGE} --debug

go-debug-fast:
	go run ${PACKAGE} --debug

go-build: go-prepare
	CGO_ENABLED=0 \
	go build \
		-ldflags '-d -s -w -extldflags=-static' \
		-tags=netgo,osusergo,static_build \
		-buildvcs=false \
		-trimpath \
		-o ${PROJECT} \
		${PACKAGE}
