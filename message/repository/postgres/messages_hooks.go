package postgres

import "chat/models"

func (u *MessageRepository) BeforeUpdate(message *models.Message) error {
	var result models.Message

	tx := u.db.Postrgres.Find(&result, message.ID)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}
