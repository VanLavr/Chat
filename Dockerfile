FROM golang:latest

WORKDIR /app

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
