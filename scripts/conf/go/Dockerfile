FROM golang:1.12.7-stretch

MAINTAINER thomas@leroux.io

ENV DEBIAN_FRONTEND noninteractive
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

RUN apt-get -y update \
    && apt-get upgrade -y \
    && apt-get -y install git \
    && go get -u github.com/stretchr/testify/require \
    && go get -u github.com/golangci/golangci-lint/cmd/golangci-lint \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && useradd -ms /bin/bash gopher

COPY . /media/novln/soba
WORKDIR /media/novln/soba
RUN chown -R gopher:gopher /media/novln/soba
ENV GOPATH /home/gopher/go
USER gopher

RUN go mod download

CMD [ "/bin/bash" ]
