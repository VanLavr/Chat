package main

import (
	"chat/internal/chatroom/repository/postgres"
	schema "chat/migrations"
	"chat/models"
	"log"
)

func main() {
	var initer = schema.NewStorage()
	initer.MigrateAll()

	r := postgres.NewChatroomRepository(schema.NewStorage())
	r.Store(&models.Chatroom{Name: "BOBBY's chat", Password: "123", CreatorID: 1})
	r.Store(&models.Chatroom{Name: "BOBBY'sdsadas chat", Password: "123", CreatorID: 1})
	if err := r.Delete(1, 2); err != nil {
		log.Println(err)
	}
}
