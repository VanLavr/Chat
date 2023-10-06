package main

import (
	"chat/internal/repository/migrations/schema"
)

func main() {
	var initer = new(schema.Database)
	initer.MigrateAll()
}
