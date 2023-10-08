package postgres

import (
	schema "chat/migrations"
	"chat/models"
)

func (u *UserRepository) BeforeAddUserToChatroom(uid, chatId int) (err error) {
	res := u.db.Postrgres.Find(&schema.UserChat{UserID: uid, ChatroomID: chatId})
	if res.RowsAffected != 0 {
		return models.ErrUserAlreadyInChat
	}

	return nil
}

func (u *UserRepository) BeforeUpdate(user *models.User) error {
	var result models.User
	tx := u.db.Postrgres.Find(&result, user.ID)
	if tx.Error != nil {
		return tx.Error
	}

	if result.ID == 0 {
		return models.ErrNotFound
	}

	return nil
}
