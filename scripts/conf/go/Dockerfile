FROM golang:1.12.1-stretch

MAINTAINER thomas@leroux.io

ENV DEBIAN_FRONTEND noninteractive
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

RUN apt-get -y update \
    && apt-get upgrade -y \
    && apt-get -y install git \
    && go get -u -v github.com/stretchr/testify/require \
    && go get -u github.com/golangci/golangci-lint/cmd/golangci-lint \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY . /media/novln/soba
WORKDIR /media/novln/soba
RUN go mod download

CMD /bin/bash
