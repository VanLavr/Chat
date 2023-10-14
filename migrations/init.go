package schema

import (
	"chat/models"
	"chat/pkg/config"
	"chat/pkg/logger"

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
		logger.STDLogger.Fatal(err.Error())
	}

	return s
}

func (a Storage) MigrateAll() {
	err := a.Postrgres.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		logger.STDLogger.Fatal(err.Error())
	}
	err = a.Postrgres.Migrator().AutoMigrate(&models.Chatroom{})
	if err != nil {
		logger.STDLogger.Fatal(err.Error())
	}
	err = a.Postrgres.Migrator().AutoMigrate(&models.Message{})
	if err != nil {
		logger.STDLogger.Fatal(err.Error())
	}
	err = a.Postrgres.Migrator().AutoMigrate(&UserChat{})
	if err != nil {
		logger.STDLogger.Fatal(err.Error())
	}
}
