package main

import schema "chat/migrations"

func main() {
	var initer = new(schema.Database)
	initer.MigrateAll()
}
