FROM golang:1.21.1-bullseye AS build
#FROM kitanoyoru/effective-mobile/base-builder:latest AS build

LABEL maintainer="Alexandr Rutkowski <kitanoyoru@protonmail.com>"

ENV PROJECT_DIR = "/usr/src/app"

WORKDIR ${PROJECT_DIR}

ADD go.mod .
ADD go.sum .

COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN go mod tidy
RUN go build -o /bin/app cmd/main.go
RUN rm -rf ${PROJECT_DIR}

# REFACTOR: chnage to debian-slim
FROM ubuntu:latest

ENV USER kitanoyoru

RUN useradd -m -U -d /home/${USER} ${USER} -s /bin/bash

RUN set -ex; \
  BUILD_DEPS='ca-certificates vim'; \
  PREFIX=/usr/local; \
  apt-get update; \
  apt-get install -y $BUILD_DEPS --no-install-recommends;

RUN mkdir /var/run/app
RUN chown ${USER}:${USER} /var/run/app

COPY --from=build /bin/app /usr/local/bin/app

USER ${USER}
WORKDIR /home/${USER}

CMD ["app", "server"]
