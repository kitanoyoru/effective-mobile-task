FROM golang:1.21.1-bullseye AS build

LABEL maintainer="Alexandr Rutkowski <kitanoyoru@protonmail.com>"

ENV PROJECT_DIR = "/usr/src/app"

WORKDIR ${PROJECT_DIR}

ADD go.mod .
ADD go.sum .

COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN go mod tidy
RUN go build -o /bin/effective-mobile-task cmd/main.go
RUN rm -rf ${PROJECT_DIR}

FROM debian:buster-slim

ENV USER kitanoyoru

RUN useradd -m -U -d /home/${USER} ${USER} -s /bin/bash

RUN set -ex; \
  BUILD_DEPS='ca-certificates vim'; \
  PREFIX=/usr/local; \
  apt-get update; \
  apt-get install -y $BUILD_DEPS --no-install-recommends;

RUN mkdir /var/run/app
RUN chown ${USER}:${USER} /var/run/app


COPY --from=build /bin/effective-mobile-task /usr/local/bin/effective-mobile-task

COPY static /usr/share/effective-mobile-task

COPY infra/docker-entrypoint.sh /usr/local/bin/
RUN chmod 755 /usr/local/bin/docker-entrypoint.sh


USER ${USER}
WORKDIR /home/${USER}

ENTRYPOINT ["docker-entrypoint.sh"]
