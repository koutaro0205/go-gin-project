FROM golang:1.22.2-alpine

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

RUN go mod download
