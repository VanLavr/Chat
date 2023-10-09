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

	return result, nil
}

func (c *ChatroomRepository) Store(Chatroom *models.Chatroom) error {
	if err := c.beforeCreate(Chatroom); err != nil {
		return err
	}

	tx := c.db.Postrgres.Save(Chatroom)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (c *ChatroomRepository) Update(Chatroom *models.Chatroom) error {
	if err := c.beforeUpdate(Chatroom); err != nil {
		return err
	}

	tx := c.db.Postrgres.Save(Chatroom)
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
