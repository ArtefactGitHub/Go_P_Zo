# syntax = docker/dockerfile:1.2

FROM golang:1.16-alpine

ARG Go_P_Zo_DB_USER=testuser
ARG Go_P_Zo_DB_PASSWORD
ARG Go_P_Zo_DB_ADDRESS=127.0.0.1:13306
ARG Go_P_Zo_AUTH_SIGNKEY

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
WORKDIR cmd
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 8000

CMD ["app"]
