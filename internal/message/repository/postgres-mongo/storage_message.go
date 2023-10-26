package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"chat/pkg/logger"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type messageRepository struct {
	db *schema.Storage
}

func NewMessageRepository(db *schema.Storage) models.MessageRepository {
	return &messageRepository{db: db}
}

func (m *messageRepository) Fetch(limit int) ([]models.Message, error) {
	var result []models.Message
	if limit == 0 {
		tx := m.db.Postrgres.Find(&result)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := m.db.Postrgres.Limit(limit).Find(&result)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return result, nil
}

func (m *messageRepository) FetchOne(id int) (models.Message, error) {
	var result models.Message
	tx := m.db.Postrgres.Where("id = ?", id).Find(&result)
	if tx.Error != nil {
		return models.Message{}, tx.Error
	}

	return result, nil
}

func (m *messageRepository) Store(Message models.Message) error {
	if err := m.beforeCreate(Message); err != nil {
		return err
	}
	tx := m.db.Postrgres.Save(&Message)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *messageRepository) Update(Message models.Message) error {
	if err := m.beforeUpdate(Message); err != nil {
		return err
	}

	tx := m.db.Postrgres.Save(&Message)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *messageRepository) Delete(id int) error {
	if err := m.beforeDelete(id); err != nil {
		return err
	}
	tx := m.db.Postrgres.Delete(&models.Message{ID: id})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *messageRepository) FetchByUserID(limit, id int) ([]models.Message, error) {
	var result []models.Message
	tx := m.db.Postrgres.Where("user_id = ?", id).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (m *messageRepository) FetchByChatroomID(limit, id int) ([]models.Message, error) {
	var result []models.Message
	tx := m.db.Postrgres.Where("chatroom_id = ?", id).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (m *messageRepository) StorePhoto(Message models.Message) (string, error) {
	imagesCollection := m.db.SetMongoOnStatic()

	dao, err := m.castToDAO(Message)
	if err != nil {
		return "", err
	}

	result, err := imagesCollection.InsertOne(context.TODO(), dao)
	if err != nil {
		logger.STDLogger.Info(err.Error())
		logger.FileLogger.Info(err.Error())
		return "", models.ErrInternalServerError
	}

	id := result.InsertedID
	log.Println(id)
	oid, ok := id.(primitive.ObjectID)
	if !ok {
		logger.STDLogger.Error("can not assert object id")
		logger.FileLogger.Error("can not assert object id")

		return "", models.ErrInternalServerError
	}

	runed := []rune(oid.String())
	runed = append(runed[:len(runed)-1], runed[len(runed):]...)
	runed = append(runed[:len(runed)-1], runed[len(runed):]...)
	runed = runed[len("ObjectID(\""):]

	return string(runed), nil
}

func (m *messageRepository) DeletePhoto(id string) (int64, error) {
	imagesCollection := m.db.SetMongoOnStatic()
	documentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": documentID}

	result, err := imagesCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		logger.STDLogger.Error(err.Error())
		logger.FileLogger.Error(err.Error())

		return 0, err
	}

	return result.DeletedCount, nil
}

func (m *messageRepository) FindPhoto(message models.Message) (string, error) {
	imagesCollection := m.db.SetMongoOnStatic()
	dao, err := m.castToDAO(message)
	if err != nil {
		return "", err
	}

	var document map[string]interface{}
	err = imagesCollection.FindOne(context.TODO(), bson.M{"userid": dao.UserID, "chatroomid": dao.ChatroomID, "timestamp": dao.Timestamp}).Decode(&document)
	if err != nil {
		return "", err
	}

	id := document["_id"]
	log.Println(id)
	oid, ok := id.(primitive.ObjectID)
	if !ok {
		logger.STDLogger.Error("can not assert object id")
		logger.FileLogger.Error("can not assert object id")

		return "", models.ErrInternalServerError
	}

	runed := []rune(oid.String())
	runed = append(runed[:len(runed)-1], runed[len(runed):]...)
	runed = append(runed[:len(runed)-1], runed[len(runed):]...)
	runed = runed[len("ObjectID(\""):]

	log.Println(string(runed))
	return string(runed), nil
}

// message.Sended.IsZero()
func (m *messageRepository) castToDAO(message models.Message) (*ImageDAO, error) {
	if message.ChatroomID == 0 || message.UserID == 0 {
		return nil, models.ErrBadParamInput
	}

	return &ImageDAO{
		UserID:     message.UserID,
		ChatroomID: message.ChatroomID,
		Timestamp:  message.Sended.Format(models.TimeLayout),
	}, nil
}
