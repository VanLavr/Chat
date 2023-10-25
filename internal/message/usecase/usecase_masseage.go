package usecase

import (
	"chat/models"
	"chat/pkg/logger"
	"errors"
)

type usecase struct {
	repo models.MessageRepository
}

func NewUsecase(repo models.MessageRepository) models.MessageUsecase {
	return &usecase{repo: repo}
}

func (u *usecase) GetChatMessages(limit, id int) ([]models.Message, error) {
	messages, err := u.repo.FetchByChatroomID(limit, id)
	if err != nil {
		logger.FileLogger.Fatal(err.Error())
	}

	return messages, nil
}

func (u *usecase) GetUserMessages(limit, id int) ([]models.Message, error) {
	messages, err := u.repo.FetchByUserID(limit, id)
	if err != nil {
		logger.FileLogger.Fatal(err.Error())
	}

	return messages, nil
}

func (u *usecase) GetMessages(limit int) ([]models.Message, error) {
	messages, err := u.repo.Fetch(limit)
	if err != nil {
		logger.FileLogger.Fatal(err.Error())
	}

	return messages, nil
}

func (u *usecase) GetById(id int) (models.Message, error) {
	message, err := u.repo.FetchOne(id)
	if err != nil {
		logger.FileLogger.Fatal(err.Error())
	}

	return message, nil
}

func (u *usecase) CreateMessage(message models.Message) error {
	if err := u.repo.Store(message); err != nil {
		if errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrEmptyFields) {
			return err
		} else {
			logger.FileLogger.Fatal(err.Error())
		}
	}

	return nil
}

func (u *usecase) UpdateMessage(message models.Message) error {
	if err := u.repo.Update(message); err != nil {
		if errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrEmptyFields) {
			return err
		} else {
			logger.FileLogger.Fatal(err.Error())
		}
	}

	return nil
}

func (u *usecase) DeleteMessage(id int) error {
	if err := u.repo.Delete(id); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return err
		} else {
			logger.FileLogger.Fatal(err.Error())
		}
	}

	return nil
}

func (u *usecase) StorePhoto(message models.Message) (string, error) {
	if message.ChatroomID == 0 || message.UserID == 0 {
		return "", models.ErrBadParamInput
	}

	id, err := u.repo.StorePhoto(message)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *usecase) DeletePhoto(id string) (int64, error) {
	if id == "" {
		return 0, models.ErrBadParamInput
	}

	deleted, err := u.repo.DeletePhoto(id)
	if err != nil {
		return 0, models.ErrInternalServerError
	}

	return deleted, nil
}
