package usecase

import (
	"chat/models"
	"errors"
	"log"
)

type usecase struct {
	repo models.ChatroomRepository
}

// GetById(id int) Chatroom
// Get(limit int) []Chatroom
// CreateChat(chatroom Chatroom) error
// DeleteChat(chat Chatroom) error
// UpdateChat(chat Chatroom) error

func NewUsecase(repo models.ChatroomRepository) models.ChatroomUsecase {
	return &usecase{repo: repo}
}

func (u *usecase) GetById(id int) (models.Chatroom, error) {
	chat, err := u.repo.FetchOne(id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.Chatroom{}, err
		} else {
			log.Fatal(err)
		}
	}

	return chat, nil
}

func (u *usecase) Get(limit int) ([]models.Chatroom, error) {
	chats, err := u.repo.Fetch(limit)
	if err != nil {
		log.Fatal(err)
	}

	return chats, nil
}

func (u *usecase) CreateChat(chatroom models.Chatroom) error {
	if err := u.repo.Store(&chatroom); err != nil {
		if errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrAlreadyExists) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (u *usecase) DeleteChat(chat models.Chatroom) error {
	if err := u.repo.Delete(chat.ID); err != nil {
		if errors.Is(err, models.ErrPermisionDenied) || errors.Is(err, models.ErrNotFound) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (u *usecase) UpdateChat(chat models.Chatroom) error {
	if err := u.repo.Update(&chat); err != nil {
		if errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrNotFound) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}
