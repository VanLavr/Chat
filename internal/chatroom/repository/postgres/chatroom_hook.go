package postgres

import "chat/models"

func (c *ChatroomRepository) beforeUpdate(Chatroom *models.Chatroom) error {
	if Chatroom.Name == "" || Chatroom.Password == "" {
		return models.ErrEmptyFields
	}

	var result models.Chatroom
	tx := c.db.Postrgres.Where(Chatroom.ID).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (c *ChatroomRepository) beforeCreate(Chatroom *models.Chatroom) error {
	if Chatroom.Name == "" || Chatroom.Password == "" {
		return models.ErrEmptyFields
	}

	var result models.Chatroom
	tx := c.db.Postrgres.Where("name = ?", Chatroom.Name).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID != 0 {
		return models.ErrAlreadyExists
	}

	return nil
}

func (c *ChatroomRepository) beforeDelete(id int) error {
	var result models.Chatroom
	tx := c.db.Postrgres.Where(id).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}
