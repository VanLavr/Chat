package postgres

import (
	"chat/models"
)

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
