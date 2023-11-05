package schema

import (
	"context"
	"log"

	"chat/models"
	"chat/pkg/config"
	"chat/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
}

type storage struct {
	Postrgres *gorm.DB
	Mongo     *mongo.Client
}

func NewStorage() Storage {
	s := new(storage)
	dsn := config.Con.GetPostgres()
	log.Println(dsn)

	var err error
	s.Postrgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.STDLogger.Fatal(err.Error())
	}

	mongoURI := config.Con.GetMongoHost()
	s.Mongo, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		logger.STDLogger.Fatal(err.Error())
	}

	return s
}

func (a storage) MigrateAll() {
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

func (s storage) SetMongoOnStatic() *mongo.Collection {
	return s.Mongo.Database("chatserver_static").Collection("images")
}
