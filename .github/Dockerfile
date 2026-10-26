# syntax = docker/dockerfile:1

# build backend
FROM golang:latest AS build-go
WORKDIR /build

COPY . /build
RUN \
	--mount=type=cache,target=/root/.cache \
	--mount=type=cache,target=/go \
	make go-build

# build frontend
FROM oven/bun:1-alpine AS build-web

# nodejs is needed for "nuxt prepare" (postinstall), as bun doesn't implement
# node:module.findPackageJSON, which nuxt modules (e.g. nuxt-og-image) rely on.
RUN apk add --no-cache make bash curl nodejs

WORKDIR /build
COPY . /build
RUN make node-prepare
RUN make node-build

# runtime
FROM oven/bun:1-alpine AS runtime
WORKDIR /app

RUN apk add --no-cache ca-certificates

# copy frontend
COPY --from=build-web /build/web/.output /app/web

# copy backend
COPY --from=build-go /build/httpserver /app/backend/httpserver

ENTRYPOINT ["/bin/ash", "-i"]
