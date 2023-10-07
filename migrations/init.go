package schema

import (
	"chat/models"
	"chat/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Postrgres *gorm.DB
}

func NewStorage() *Storage {
	s := new(Storage)
	dsn := config.Con.GetPostgres()

	var err error
	s.Postrgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return s
}

func (a Storage) MigrateAll() {
	err := a.Postrgres.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
	err = a.Postrgres.Migrator().AutoMigrate(&models.Chatroom{})
	if err != nil {
		log.Fatal(err)
	}
	err = a.Postrgres.Migrator().AutoMigrate(&models.Message{})
	if err != nil {
		log.Fatal(err)
	}
	err = a.Postrgres.Migrator().AutoMigrate(&UserChat{})
	if err != nil {
		log.Fatal(err)
	}
}
