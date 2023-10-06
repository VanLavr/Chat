package schema

import (
	"chat/internal/service/models"
	"chat/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Postrgres *gorm.DB
}

func (d Database) MigrateAll() {
	dsn := config.Con.GetPostgres()

	var err error
	d.Postrgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = d.Postrgres.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	err = d.Postrgres.Migrator().AutoMigrate(&models.ChatRoom{})
	if err != nil {
		log.Fatal(err)
	}
	err = d.Postrgres.Migrator().AutoMigrate(&models.Message{})
	if err != nil {
		log.Fatal(err)
	}
	err = d.Postrgres.Migrator().AutoMigrate(&models.UserChat{})
	if err != nil {
		log.Fatal(err)
	}
}
