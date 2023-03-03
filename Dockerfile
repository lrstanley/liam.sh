# syntax = docker/dockerfile:1.4

# frontend
FROM node:19 as build-node
WORKDIR /build

COPY . /build
RUN \
	--mount=type=cache,target=/build/cmd/httpserver/public/node_modules \
	make node-fetch node-build

# backend
FROM golang:latest as build-go
WORKDIR /build

COPY . /build
COPY --from=build-node /build/cmd/httpserver/public/dist/ /build/cmd/httpserver/public/dist/
RUN \
	--mount=type=cache,target=/root/.cache \
	--mount=type=cache,target=/go \
	make go-build

# runtime
FROM alpine:3.17
WORKDIR /app

RUN apk add --no-cache ca-certificates
RUN if [ ! -e /etc/nsswitch.conf ];then echo 'hosts: files dns' > /etc/nsswitch.conf;fi
COPY --from=build-go /build/httpserver /app/httpserver

VOLUME /app/.gitapicache
CMD ["/app/httpserver"]
