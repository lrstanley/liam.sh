.DEFAULT_GOAL := build-all

export PROJECT := "httpserver"
export PACKAGE := "github.com/lrstanley/liam.sh/cmd/httpserver"
export DOCKER_BUILDKIT := 1
export LICENSE_IGNORE := "graphql-tag"
export KUBERNETES_NAMESPACE := "liam-sh"
export KUBERNETES_SELECTOR := "app.kubernetes.io/name=liam-sh"

license:
	curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

prepare: clean node-prepare node-build go-prepare
	@echo

build-all: prepare go-build
	@echo

up: node-upgrade-deps go-upgrade-deps
	@echo

clean:
	/bin/rm -rfv "cmd/httpserver/public/dist/*" ${PROJECT}

mirror-prod-db:
	rm -rfv local.db*
	kubectl -n ${KUBERNETES_NAMESPACE} get pods -l ${KUBERNETES_SELECTOR} --no-headers \
		| awk '{print $$1}' | head -1 \
		| xargs -n1 -I{} kubectl cp -n ${KUBERNETES_NAMESPACE} {}:/data/local.db local.db

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
		--tag ${PROJECT} \
		--pull \
		--force-rm .

# frontend
node-fetch:
	command -v pnpm >/dev/null >&2 || npm install \
		--no-audit \
		--no-fund \
		--quiet \
		--global pnpm
	cd cmd/httpserver/public && \
		pnpm install

node-upgrade-deps:
	cd cmd/httpserver/public && \
		pnpm up -iL

node-prepare: license node-fetch
	cd cmd/httpserver/public && \
		pnpm exec prettier \
			--cache \
			--write \
			--ignore-path src/lib/api/graphql.ts \
			src/

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
		-installsuffix netgo \
		-buildvcs=false \
		-trimpath \
		-o ${PROJECT} \
		${PACKAGE}
