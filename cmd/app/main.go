package main

import schema "chat/migrations"

func main() {
	var initer = schema.NewStorage()
	initer.MigrateAll()
}
