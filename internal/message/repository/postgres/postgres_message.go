package postgres

import (
	schema "chat/migrations"
	"chat/models"
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
