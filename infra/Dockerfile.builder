FROM alpine:latest

ENV GO_VERSION=1.21.1
ENV GOPATH=/go
ENV PATH=${PATH}:${GOPATH}/bin

RUN apk update && apk add --no-cache wget curl git

RUN wget -q https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz

ENV GOROOT=/usr/local/go
