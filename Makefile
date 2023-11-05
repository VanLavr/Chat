run:
	go mod tidy
	go run ./cmd/app/main.go

build_linux:
	go build -o ./bin/main ./cmd/app/main.go

build_windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/main_win ./cmd/app/main.go

test:
	echo "REPO:PG"
	go test ./internal/user/repository/postgres/
	go test ./internal/chatroom/repository/postgres/
	go test ./internal/message/repository/postgres/

Create_image:
	sudo docker buildx build . --tag lavrushkoivan/web_chat

up_container:
	make Create_image
	sudo docker-compose up -d