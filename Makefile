run:
	go mod tidy
	go run ./cmd/app/main.go

build_linux:
	go build -o ./bin/main ./cmd/app/main.go

build_windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/main_win ./cmd/app/main.go

test:
	echo "REPO:PG"
	go test ./user/repository/postgres/
	go test ./chatroom/repository/postgres/