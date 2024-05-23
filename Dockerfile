FROM golang:alpine3.20 as builder

WORKDIR /exercism

COPY . .

RUN set -xe \
    && apk update \
    && apk upgrade \
    && apk add --update --no-cache build-base bash


ENTRYPOINT [ "/bin/sh" ]
CMD [ "-c", "while true; do echo 'container is running...' && sleep 300; done" ]
