FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN go build -o ./bin/main ./cmd/app/main.go

EXPOSE 8080

CMD ["./bin/main"]
