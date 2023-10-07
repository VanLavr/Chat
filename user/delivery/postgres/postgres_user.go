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
