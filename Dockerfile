FROM golang:1.21.1-alpine3.18

LABEL authors="Jasmeet Singh"
LABEL maintainer="thejasmeet.aws@gmail.com"

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN mkdir /code
WORKDIR /code

COPY . /code/