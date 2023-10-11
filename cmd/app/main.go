package main

import (
	schema "chat/migrations"
)

func main() {
	var initer = schema.NewStorage()
	initer.MigrateAll()

	// ur := postgres.NewUserRepository(schema.NewStorage())
	// if err := ur.AddUserToChatroom(2, 1); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("ok")
	// }
}
