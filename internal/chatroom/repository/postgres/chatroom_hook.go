package postgres

import (
	"chat/models"
	"fmt"
	"log"
)

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
		log.Println(err)
		return err
	}

	fmt.Println(author)
	author.RoomsOwned++
	fmt.Println(author)
	if err := c.db.Postrgres.Save(&author).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *ChatroomRepository) beforeDelete(deleter, id int) error {
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

	fmt.Println(author, result, deleter)
	if !author.IsAdmin || deleter != result.CreatorID {
		return models.ErrPermisionDenied
	}

	return nil
}
