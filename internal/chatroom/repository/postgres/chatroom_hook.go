package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"chat/pkg/logger"
	"fmt"
)

func (c *chatroomRepository) beforeUpdate(Chatroom models.Chatroom) error {
	if Chatroom.Name == "" || Chatroom.Password == "" {
		return models.ErrEmptyFields
	}

	var result models.Chatroom
	tx := c.db.Postrgres.Where("id = ?", Chatroom.ID).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (c *chatroomRepository) beforeCreate(Chatroom models.Chatroom) error {
	if Chatroom.Name == "" || Chatroom.Password == "" || Chatroom.CreatorID == 0 {
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

	var author models.User
	if err := c.db.Postrgres.Where("id = ?", Chatroom.CreatorID).Find(&author).Error; err != nil {
		logger.STDLogger.Warn(err.Error())
		return err
	}

	if author.ID != 0 {
		author.RoomsOwned++
		if err := c.db.Postrgres.Save(&author).Error; err != nil {
			logger.STDLogger.Warn(err.Error())
			return err
		}
	} else {
		return models.ErrNotFound
	}

	return nil
}

func (c *chatroomRepository) beforeDelete(deleter, id int) error {
	var result models.Chatroom
	tx := c.db.Postrgres.Where(id).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	var author models.User
	if err := c.db.Postrgres.Where("id = ?", result.CreatorID).Find(&author).Error; err != nil {
		return err
	}

	var deleterUser models.User
	if err := c.db.Postrgres.Where("id = ?", deleter).Find(&deleterUser).Error; err != nil {
		return err
	}

	if !(deleterUser.ID == author.ID || deleterUser.IsAdmin) {
		return models.ErrPermisionDenied
	}

	if author.ID != 0 {
		author.RoomsOwned--
		if err := c.db.Postrgres.Save(&author).Error; err != nil {
			logger.STDLogger.Warn(err.Error())
			return err
		}
	}

	return nil
}

func (c *chatroomRepository) beforeAddUserToChatroom(uid, chatId int) (err error) {
	var user models.User
	var chat models.Chatroom
	c.db.Postrgres.Where("id = ?", uid).Find(&user)
	c.db.Postrgres.Where("id = ?", chatId).Find(&chat)
	if user.ID == 0 || chat.ID == 0 {
		return models.ErrNotFound
	}

	var uc schema.UserChat
	if err := c.db.Postrgres.Where("user_id = ?", uid).Where("chatroom_id = ?", chatId).Find(&uc).Error; err != nil {
		logger.STDLogger.Fatal(err.Error())
	}
	fmt.Println(uc)
	if uc.ChatroomID != 0 && uc.UserID != 0 {
		return models.ErrUserAlreadyInChat
	}

	return nil
}
