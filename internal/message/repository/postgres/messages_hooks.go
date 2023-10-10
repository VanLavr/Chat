package postgres

import "chat/models"

func (u *MessageRepository) beforeUpdate(message models.Message) error {
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

func (c *MessageRepository) beforeCreate(Message models.Message) error {
	if Message.Content == "" {
		return models.ErrEmptyFields
	}

	var user models.User
	var chat models.Chatroom
	if err := c.db.Postrgres.Where("user_id = ?", Message.UserID).Find(&user).Error; err != nil {
		return err
	}
	if err := c.db.Postrgres.Where("chatroom_id = ?", Message.ChatroomID).Find(&chat).Error; err != nil {
		return err
	}

	if user.ID != 0 || chat.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (c *MessageRepository) beforeDelete(id int) error {
	var result models.Message
	tx := c.db.Postrgres.Where(id).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}
