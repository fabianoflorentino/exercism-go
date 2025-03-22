FROM golang:alpine3.21 as builder

WORKDIR /exercism

COPY . .

RUN set -xe \
    && apk update \
    && apk upgrade \
    && go mod download \
    && apk add --update --no-cache build-base bash


ENTRYPOINT [ "/bin/sh" ]
CMD [ "-c", "while true; do echo 'OK! Im running...' && sleep 300; done" ]
