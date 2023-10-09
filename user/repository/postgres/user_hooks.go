package postgres

import (
	schema "chat/migrations"
	"chat/models"
)

func (u *UserRepository) beforeAddUserToChatroom(uid, chatId int) (err error) {
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
	tx := u.db.Postrgres.Where(&models.User{Name: user.Name}).Find(&result)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID != 0 {
		return models.ErrAlreadyExists
	}

	return nil
}
