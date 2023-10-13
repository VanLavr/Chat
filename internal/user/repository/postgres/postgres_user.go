package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"errors"
	"log"
)

type userRepository struct {
	db *schema.Storage
}

// Fetch(limit int) ([]User, error)
// FetchOne(id int) (User, error)
// FetchFewCertain(id ...int) ([]User, error)
// AddUserToChatroom(uid, chatId int) error
// RemoveUserFromChatroom(uid, chatId int) error
// Store(user User) error
// Update(user User) error
// Delete(id int) error

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

func (u *userRepository) GetChatters() []models.User {

	var chatusers []schema.UserChat

	if err := u.db.Postrgres.Find(&chatusers).Error; err != nil {
		log.Fatal(err)
	}

	var ids []int
	for _, id := range chatusers {
		ids = append(ids, id.UserID)
	}

	var chatters []models.User
	if err := u.db.Postrgres.Find(&chatters, ids).Error; err != nil {
		log.Fatal(err)
	}

	return chatters
}

func (u *userRepository) GetUserPassword(id int) (string, error) {
	err := u.beforeDelete(id)
	if err != nil && errors.Is(err, models.ErrNotFound) {
		return "", err
	} else if err != nil && !errors.Is(err, models.ErrNotFound) {
		log.Fatal(err)
	}

	var user models.User
	if err := u.db.Postrgres.Where("id = ?", id).Find(&user).Error; err != nil {
		log.Fatal(err)
	}

	return user.Password, nil
}
