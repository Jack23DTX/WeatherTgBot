FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY .env ./
COPY go.mod go.sum ./

RUN go mod download

COPY entity ./entity
COPY weather ./weather
COPY cmd ./cmd

RUN go build -o /app/projectbot1 ./cmd/

CMD ["/app/projectbot1"]