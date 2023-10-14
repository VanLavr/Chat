package usecase

import (
	"chat/models"
	"chat/pkg/hash"
	"errors"
	"log"
)

type usecase struct {
	repo models.UserRepository
}

func NewUsecase(repo models.UserRepository) models.UserUsecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) GetById(uid int) (models.User, error) {
	user, err := u.repo.FetchOne(uid)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.User{}, err
		} else {
			log.Fatal(err)
		}
	}

	return user, nil
}

func (u *usecase) GetUsers(limit int) []models.User {
	users, err := u.repo.Fetch(limit)
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func (u *usecase) CreateUser(user models.User) error {
	hashed := hash.Hshr.Hash(user.Password)
	user.Password = hashed

	if err := u.repo.Store(user); err != nil {
		if errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrAlreadyExists) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (u *usecase) DeleteUser(id int) error {
	if err := u.repo.Delete(id); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return err
		}
	}

	return nil
}

func (u *usecase) UpdateUser(user models.User) error {
	if err := u.repo.Update(user); err != nil {
		if errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrNotFound) {
			return err
		}
	}

	return nil
}

func (u *usecase) MakeHub() []models.User {
	return u.repo.GetChatters()
}

func (u *usecase) ValidatePassword(uid int, password string) (bool, error) {
	pwd, err := u.repo.GetUserPassword(uid)
	if err != nil {
		return false, models.ErrNotFound
	}

	if !hash.Hshr.Validate(pwd, password) {
		return false, models.ErrPermisionDenied
	}

	return true, nil
}

func (u *usecase) ValidateIncommer(uid, cid int) bool {
	return u.repo.BeforeJoin(uid, cid)
}
