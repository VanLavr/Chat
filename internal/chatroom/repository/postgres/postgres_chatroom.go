package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"log"
)

type chatroomRepository struct {
	db *schema.Storage
}

func NewChatroomRepository(db *schema.Storage) models.ChatroomRepository {
	return &chatroomRepository{db: db}
}

func (c *chatroomRepository) Fetch(limit int) ([]models.Chatroom, error) {
	var result []models.Chatroom
	if limit == 0 {
		tx := c.db.Postrgres.Find(&result)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := c.db.Postrgres.Limit(limit).Find(&result)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return result, nil
}

func (c *chatroomRepository) FetchOne(id int) (models.Chatroom, error) {
	var result models.Chatroom
	tx := c.db.Postrgres.First(&result, id)
	if tx.Error != nil {
		return models.Chatroom{}, tx.Error
	}

	if result.ID == 0 {
		return models.Chatroom{}, models.ErrNotFound
	}

	return result, nil
}

func (c *chatroomRepository) Store(chat models.Chatroom) error {
	if err := c.beforeCreate(chat); err != nil {
		return err
	}

	tx := c.db.Postrgres.Save(chat)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (c *chatroomRepository) Update(chat models.Chatroom) error {
	if err := c.beforeUpdate(chat); err != nil {
		return err
	}

	tx := c.db.Postrgres.Save(chat)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (c *chatroomRepository) Delete(deleter, id int) error {
	if err := c.beforeDelete(deleter, id); err != nil {
		return err
	}

	tx := c.db.Postrgres.Delete(&models.Chatroom{ID: id})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (c *chatroomRepository) GetRoomPassword(id int) (string, error) {
	var result models.Chatroom
	tx := c.db.Postrgres.Where("id = ?", id).Find(&result)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	if result.ID == 0 {
		return "", models.ErrNotFound
	}

	var chat models.Chatroom
	if err := c.db.Postrgres.Where("id = ?", id).Find(&chat).Error; err != nil {
		log.Fatal(err)
	}

	return chat.Password, nil
}

func (c *chatroomRepository) AddUserToChatroom(uid, chatId int) error {
	if err := c.beforeAddUserToChatroom(uid, chatId); err != nil {
		return err
	}

	tx := c.db.Postrgres.Save(&schema.UserChat{UserID: uid, ChatroomID: chatId})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (c *chatroomRepository) RemoveUserFromChatroom(uid, chatId int) error {
	tx := c.db.Postrgres.Where("user_id = ?", uid).Where("chatroom_id = ?", chatId).Delete(&schema.UserChat{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
