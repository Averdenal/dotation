FROM golang:1.15.5-alpine3.12

WORKDIR /go/src/app
COPY . /go/src/app

VOLUME /go/src/app
EXPOSE 8080
