package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"fmt"
	"log"
)

func (u *userRepository) beforeAddUserToChatroom(uid, chatId int) (err error) {
	var user models.User
	var chat models.Chatroom
	u.db.Postrgres.Where("id = ?", uid).Find(&user)
	u.db.Postrgres.Where("id = ?", chatId).Find(&chat)
	if user.ID == 0 || chat.ID == 0 {
		return models.ErrNotFound
	}

	var uc schema.UserChat
	if err := u.db.Postrgres.Where("user_id = ?", uid).Where("chatroom_id = ?", chatId).Find(&uc).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println(uc)
	if uc.ChatroomID != 0 && uc.UserID != 0 {
		return models.ErrUserAlreadyInChat
	}

	return nil
}

func (u *userRepository) beforeUpdate(user models.User) error {
	if user.Name == "" || user.Password == "" {
		return models.ErrEmptyFields
	}

	var result models.User
	tx := u.db.Postrgres.Where("id = ?", user.ID).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (u *userRepository) beforeDelete(id int) error {
	var result models.User
	tx := u.db.Postrgres.Where("id = ?", id).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (u *userRepository) beforeCreate(user models.User) error {
	if user.Name == "" || user.Password == "" {
		return models.ErrEmptyFields
	}

	var result models.User
	tx := u.db.Postrgres.Where("name = ?", user.Name).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID != 0 {
		return models.ErrAlreadyExists
	}

	return nil
}
