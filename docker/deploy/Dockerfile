FROM golang:1.18.5-buster AS builder

ENV GO111MODULE on
ENV APP_PATH=/app

RUN apt-get update && apt-get install -y git
RUN go install github.com/cespare/reflex@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN cp /go/bin/reflex /usr/bin/reflex
RUN cp /go/bin/dlv /usr/bin/dlv

WORKDIR $APP_PATH

#COPY go.mod .
#COPY go.sum .
COPY . .

RUN go mod download

