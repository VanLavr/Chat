package usecase

import (
	"chat/models"
	"errors"
	"log"
)

type usecase struct {
	repo models.MessageRepository
}

// GetChatMessages(message Message) ([]Message, error)
// GetUserMessages(message Message) ([]Message, error)
// GetMessages(limit int) ([]Message, error)
// GetById(id int) ([]Message, error)
// CreateMessage(message Message) error
// UpdateMessage(message Message) error
// DeleteMessage(message Message) error

func NewUsecase(repo models.MessageRepository) models.MessageUsecase {
	return &usecase{repo: repo}
}

func (u *usecase) GetChatMessages(limit, id int) ([]models.Message, error) {
	messages, err := u.repo.FetchByChatroomID(limit, id)
	if err != nil {
		log.Fatal(err)
	}

	return messages, nil
}

func (u *usecase) GetUserMessages(limit, id int) ([]models.Message, error) {
	messages, err := u.repo.FetchByUserID(limit, id)
	if err != nil {
		log.Fatal(err)
	}

	return messages, nil
}

func (u *usecase) GetMessages(limit int) ([]models.Message, error) {
	messages, err := u.repo.Fetch(limit)
	if err != nil {
		log.Fatal(err)
	}

	return messages, nil
}

func (u *usecase) GetById(id int) (models.Message, error) {
	message, err := u.repo.FetchOne(id)
	if err != nil {
		log.Fatal(err)
	}

	return message, nil
}

func (u *usecase) CreateMessage(message models.Message) error {
	if err := u.repo.Store(message); err != nil {
		if errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrEmptyFields) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (u *usecase) UpdateMessage(message models.Message) error {
	if err := u.repo.Update(message); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (u *usecase) DeleteMessage(id int) error {
	if err := u.repo.Delete(id); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}
