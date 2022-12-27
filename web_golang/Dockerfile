FROM golang:1.19 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /usr/src/web_golang

COPY go.mod go.sum .

RUN go mod download && go mod verify

COPY . .