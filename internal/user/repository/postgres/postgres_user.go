package postgres

import (
	schema "chat/migrations"
	"chat/models"
)

type userRepository struct {
	db *schema.Storage
}

func NewUserRepository(db *schema.Storage) models.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Fetch(limit int) ([]models.User, error) {
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

func (u *userRepository) FetchOne(id int) (models.User, error) {
	var result models.User
	tx := u.db.Postrgres.Where("id = ?", id).First(&result)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	if result.ID == 0 {
		return models.User{}, models.ErrNotFound
	}

	return result, nil
}

func (u *userRepository) FetchFewCertain(id ...int) ([]models.User, error) {
	var result []models.User

	var ids []int
	ids = append(ids, id...)

	tx := u.db.Postrgres.Find(&result, ids)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (u *userRepository) Store(user models.User) error {
	if err := u.beforeCreate(user); err != nil {
		return err
	}

	tx := u.db.Postrgres.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *userRepository) Update(user models.User) error {
	if err := u.beforeUpdate(user); err != nil {
		return err
	}

	tx := u.db.Postrgres.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *userRepository) Delete(id int) error {
	if err := u.beforeDelete(id); err != nil {
		return err
	}

	tx := u.db.Postrgres.Where("id = ?", id).Delete(&models.User{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *userRepository) AddUserToChatroom(uid, chatId int) error {
	if err := u.beforeAddUserToChatroom(uid, chatId); err != nil {
		return err
	}

	tx := u.db.Postrgres.Save(&schema.UserChat{UserID: uid, ChatroomID: chatId})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Raw(fmt.Sprintf("delete from user_chats where user_id = %d and chatroom_id = %d", uid, chatId))
func (u *userRepository) RemoveUserFromChatroom(uid, chatId int) error {
	tx := u.db.Postrgres.Where("user_id = ?", uid).Where("chatroom_id = ?", chatId).Delete(&schema.UserChat{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
