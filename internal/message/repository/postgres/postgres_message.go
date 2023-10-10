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
	tx := m.db.Postrgres.First(&result, id)
	if tx.Error != nil {
		return models.Message{}, tx.Error
	}

	return result, nil
}

func (m *MessageRepository) Store(Message *models.Message) error {
	if err := m.beforeCreate(Message); err != nil {
		return err
	}
	tx := m.db.Postrgres.Save(Message)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (m *MessageRepository) Update(Message *models.Message) error {
	if err := m.beforeUpdate(Message); err != nil {
		return err
	}

	tx := m.db.Postrgres.Save(Message)
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
	tx := m.db.Postrgres.Where(&models.Message{UserID: id})
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (m *MessageRepository) FetchByChatroomID(limit, id int) ([]models.Message, error) {
	var result []models.Message
	tx := m.db.Postrgres.Where("id = ?", &models.Message{ChatroomID: id}).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}
