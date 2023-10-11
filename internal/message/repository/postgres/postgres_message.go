package postgres

import (
	schema "chat/migrations"
	"chat/models"
)

type MessageRepository struct {
	db *schema.Storage
}

func NewMessageRepository(db *schema.Storage) *MessageRepository {
	return &MessageRepository{db: db}
}

// Fetch(limit int) ([]Message, error)
// FetchOne(id int) (Message, error)
// FetchByUserID(limit, id int) ([]Message, error)
// FetchByChatroomID(limit, id int) ([]Message, error)
// Store(Message Message) error
// Update(Message Message) error
// Delete(id int) error

func (m *MessageRepository) Fetch(limit int) ([]models.Message, error) {
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

func (m *MessageRepository) FetchOne(id int) (models.Message, error) {
	var result models.Message
	tx := m.db.Postrgres.Where("id = ?", id).Find(&result)
	if tx.Error != nil {
		return models.Message{}, tx.Error
	}

	return result, nil
}

func (m *MessageRepository) Store(Message models.Message) error {
	if err := m.beforeCreate(Message); err != nil {
		return err
	}
	tx := m.db.Postrgres.Save(&Message)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *MessageRepository) Update(Message models.Message) error {
	if err := m.beforeUpdate(Message); err != nil {
		return err
	}

	tx := m.db.Postrgres.Save(&Message)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *MessageRepository) Delete(id int) error {
	if err := m.beforeDelete(id); err != nil {
		return err
	}
	tx := m.db.Postrgres.Delete(&models.Message{ID: id})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *MessageRepository) FetchByUserID(limit, id int) ([]models.Message, error) {
	var result []models.Message
	tx := m.db.Postrgres.Where("user_id = ?", &models.Message{UserID: id}).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (m *MessageRepository) FetchByChatroomID(limit, id int) ([]models.Message, error) {
	var result []models.Message
	tx := m.db.Postrgres.Where("chatroom_id = ?", &models.Message{ChatroomID: id}).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}
