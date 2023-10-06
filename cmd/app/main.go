package main

import (
	"chat/internal/service/migrations/schema"
)

func main() {
	var initer = new(schema.Database)
	initer.MigrateAll()
}
