FROM golang:latest as binder

COPY . /api
WORKDIR /api

ENV GO111MODULE=on

## vips
RUN apt-get update -y && \
    apt-get -y install software-properties-common &&\
    add-apt-repository -y ppa:strukturag/libde265 && \
    add-apt-repository -y ppa:strukturag/libheif && \
    add-apt-repository ppa:tonimelisma/ppa && \
    apt-get -y install libvips-dev

ENV CGO_CFLAGS_ALLOW=-Xpreprocessor

# Go Build
RUN go build -o api

ENTRYPOINT ["./api"]
