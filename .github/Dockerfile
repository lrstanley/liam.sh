# syntax = docker/dockerfile:1

# build backend
FROM golang:latest as build-go
WORKDIR /build

COPY . /build
RUN \
	--mount=type=cache,target=/root/.cache \
	--mount=type=cache,target=/go \
	make go-build

# build frontend
FROM node:22-alpine as build

RUN apk add --no-cache make bash curl

WORKDIR /build/frontend
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

COPY frontend/package.json frontend/pnpm-lock.yaml /build/frontend
COPY Makefile /build

WORKDIR /build
RUN --mount=type=cache,id=pnpm,target=/pnpm/store make node-prepare
COPY . /build
RUN \
	--mount=type=cache,id=pnpm,target=/pnpm/store \
	--mount=type=secret,id=NUXT_UI_PRO_LICENSE,env=NUXT_UI_PRO_LICENSE \
	make node-build

# runtime
FROM node:22-alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates

# copy frontend
COPY --from=build /build/frontend/.output /app/frontend

# copy backend
COPY --from=build-go /build/httpserver /app/backend/httpserver

ENTRYPOINT ["/bin/ash", "-i"]
