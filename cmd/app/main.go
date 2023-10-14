package main

import (
	chatroomDelivery "chat/internal/chatroom/delivery"
	chatroomRepo "chat/internal/chatroom/repository/postgres"
	chatroomUsecase "chat/internal/chatroom/usecase"

	messageDelivery "chat/internal/message/delivery"
	messageRepo "chat/internal/message/repository/postgres"
	messageUsecase "chat/internal/message/usecase"

	userDelivery "chat/internal/user/delivery"
	userRepo "chat/internal/user/repository/postgres"
	userUsecase "chat/internal/user/usecase"

	schema "chat/migrations"

	"chat/pkg/config"
	corsmiddleware "chat/pkg/cors_middleware"

	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	var storage = schema.NewStorage()
	storage.MigrateAll()

	var (
		addr string = config.Con.GetAddress()
		port string = config.Con.GetPort()
	)

	e := echo.New()
	corsmiddleware.NewMiddleware(e)

	userRepo := userRepo.NewUserRepository(storage)
	userUsecase := userUsecase.NewUsecase(userRepo)
	userDelivery.Register(e, userUsecase)

	chatroomRepo := chatroomRepo.NewChatroomRepository(storage)
	chatroomUsecase := chatroomUsecase.NewUsecase(chatroomRepo)
	chatroomDelivery.Register(e, chatroomUsecase)

	messageRepo := messageRepo.NewMessageRepository(storage)
	messageUsecase := messageUsecase.NewUsecase(messageRepo)
	messageDelivery.Register(e, messageUsecase)

	go func() {
		err := e.Start(addr + ":" + port)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
