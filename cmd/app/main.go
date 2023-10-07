package main

import schema "chat/migrations"

func main() {
	var initer = new(schema.Storage)
	initer.MigrateAll()
}
