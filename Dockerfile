
# build image
FROM golang:1.14 as build
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN make

FROM alpine:latest

RUN apk add --no-cache ca-certificates

# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/docker-library/golang/blob/1eb096131592bcbc90aa3b97471811c798a93573/1.14/alpine3.12/Dockerfile#L9
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

COPY --from=build /build/liam.sh /usr/local/bin/liam.sh

LABEL org.opencontainers.image.title=liam.sh
LABEL org.opencontainers.image.description="liam.sh"
LABEL org.opencontainers.image.url="https://liam.sh"
LABEL org.opencontainers.image.documentation="https://github.com/lrstanley/liam.sh"
LABEL org.opencontainers.image.licenses=MIT
LABEL org.opencontainers.image.source="https://github.com/lrstanley/liam.sh"

VOLUME /config
EXPOSE 80
WORKDIR /
ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
CMD ["liam.sh", "run", "--bind", "0.0.0.0:80", "--behind-proxy", "--post-dir", "/config/posts/"]
