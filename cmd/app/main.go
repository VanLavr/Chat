package main

import (
	chatroomDelivery "chat/internal/chatroom/delivery"
	chatroomRepo "chat/internal/chatroom/repository/postgres"
	chatroomUsecase "chat/internal/chatroom/usecase"
	"chat/models"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	messageDelivery "chat/internal/message/delivery"
	messageRepo "chat/internal/message/repository/postgres"
	messageUsecase "chat/internal/message/usecase"

	userDelivery "chat/internal/user/delivery"
	userRepo "chat/internal/user/repository/postgres"
	userUsecase "chat/internal/user/usecase"

	schema "chat/migrations"

	"chat/pkg/config"
	corsmiddleware "chat/pkg/cors_middleware"
	"chat/pkg/logger"

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
	e.GET("/", func(c echo.Context) error {
		logger.STDLogger.Info("/ping")
		return c.JSON(200, models.Response{
			Message: "ping",
		})
	})

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
			log.Fatal(err, "here was the error")
		}
	}()
	logger.STDLogger.Info("Server started...")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.STDLogger.Fatal(err.Error())
	}
}
