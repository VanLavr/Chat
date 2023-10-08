package postgres

import (
	schema "chat/migrations"
	"chat/models"
)

type UserRepository struct {
	db *schema.Storage
}

func NewUserRepository(db *schema.Storage) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Fetch(limit int) ([]models.User, error) {
	var result []models.User
	if limit == 0 {
		tx := u.db.Postrgres.Find(&result)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := u.db.Postrgres.Limit(limit).Find(&result)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return result, nil
}

func (u *UserRepository) FetchOne(id int) (models.User, error) {
	var result models.User
	tx := u.db.Postrgres.First(&result, id)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	return result, nil
}

func (u *UserRepository) FetchFewCertain(id ...int) ([]models.User, error) {
	var result []models.User
	tx := u.db.Postrgres.Find(&result, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (u *UserRepository) Store(user *models.User) error {
	tx := u.db.Postrgres.Save(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *UserRepository) Update(user *models.User) error {
	if err := u.BeforeUpdate(user); err != nil {
		return err
	}

	tx := u.db.Postrgres.Save(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *UserRepository) Delete(id int) error {
	tx := u.db.Postrgres.Delete(&models.User{ID: id})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *UserRepository) AddUserToChatroom(uid, chatId int) error {
	if err := u.BeforeAddUserToChatroom(uid, chatId); err != nil {
		return err
	}

	tx := u.db.Postrgres.Save(&schema.UserChat{UserID: uid, ChatroomID: chatId})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Raw(fmt.Sprintf("delete from user_chats where user_id = %d and chatroom_id = %d", uid, chatId))
func (u *UserRepository) RemoveUserFromChatroom(uid, chatId int) error {
	tx := u.db.Postrgres.Where("user_id = ?", uid).Where("chatroom_id = ?", chatId).Delete(&schema.UserChat{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
