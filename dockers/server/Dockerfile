FROM golang:1.15.6-alpine3.12 as builder

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/cosmtrek/air \
  && chmod +x ${GOPATH}/bin/air

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .