# syntax = docker/dockerfile:1.2

FROM golang:1.18.4-alpine3.16

ENV Go_P_Zo_DB_USER=testuser
ENV Go_P_Zo_DB_PASSWORD=""
ENV Go_P_Zo_DB_ADDRESS=localhost:13306
ENV Go_P_Zo_AUTH_SIGNKEY=""
ENV Go_P_Zo_ROOT_PATH=../
ENV HOST=""

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

RUN go install github.com/cosmtrek/air@v1.40.4

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
WORKDIR cmd
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 8080

CMD ["app"]
