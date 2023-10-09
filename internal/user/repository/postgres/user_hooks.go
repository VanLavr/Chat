package postgres

import (
	schema "chat/migrations"
	"chat/models"
)

func (u *UserRepository) beforeAddUserToChatroom(uid, chatId int) (err error) {
	var user models.User
	var chat models.Chatroom
	u.db.Postrgres.Where("id = ?", uid).Find(&user)
	u.db.Postrgres.Where("id = ?", chatId).Find(&chat)
	if user.ID == 0 || chat.ID == 0 {
		return models.ErrNotFound
	}

	res := u.db.Postrgres.Find(&schema.UserChat{UserID: uid, ChatroomID: chatId})
	if res.RowsAffected != 0 {
		return models.ErrUserAlreadyInChat
	}

	return nil
}

func (u *UserRepository) beforeUpdate(user *models.User) error {
	if user.Name == "" || user.Password == "" {
		return models.ErrEmptyFields
	}

	var result models.User
	tx := u.db.Postrgres.Where(user.ID).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (u *UserRepository) beforeDelete(id int) error {
	var result models.User
	tx := u.db.Postrgres.Where(id).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (u *UserRepository) beforeCreate(user *models.User) error {
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
