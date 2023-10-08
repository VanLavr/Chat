package postgres

import "chat/models"

func (c *ChatroomRepository) beforeUpdate(Chatroom *models.Chatroom) error {
	var result models.Chatroom
	tx := c.db.Postrgres.Find(&result, Chatroom.ID)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}
