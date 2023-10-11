package postgres

import (
	schema "chat/migrations"
	"chat/models"
)

type ChatroomRepository struct {
	db *schema.Storage
}

func NewChatroomRepository(db *schema.Storage) *ChatroomRepository {
	return &ChatroomRepository{db: db}
}

// Fetch(limit int) ([]Chatroom, error)
// FetchOne(id int) (Chatroom, error)
// Store(chat Chatroom) error
// Update(chat Chatroom) error
// Delete(deleter, id int) error

func (c *ChatroomRepository) Fetch(limit int) ([]models.Chatroom, error) {
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

func (c *ChatroomRepository) FetchOne(id int) (models.Chatroom, error) {
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

func (c *ChatroomRepository) Store(chat models.Chatroom) error {
	if err := c.beforeCreate(chat); err != nil {
		return err
	}

	tx := c.db.Postrgres.Save(chat)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (c *ChatroomRepository) Update(chat models.Chatroom) error {
	if err := c.beforeUpdate(chat); err != nil {
		return err
	}

	tx := c.db.Postrgres.Save(chat)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (c *ChatroomRepository) Delete(deleter, id int) error {
	if err := c.beforeDelete(deleter, id); err != nil {
		return err
	}

	tx := c.db.Postrgres.Delete(&models.Chatroom{ID: id})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
