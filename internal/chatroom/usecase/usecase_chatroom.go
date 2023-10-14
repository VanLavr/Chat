package usecase

import (
	"chat/models"
	"chat/pkg/hash"
	"errors"
	"log"
)

type usecase struct {
	repo models.ChatroomRepository
}

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
	hashed := hash.Hshr.Hash(chatroom.Password)
	chatroom.Password = hashed

	if err := u.repo.Store(chatroom); err != nil {
		if errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrAlreadyExists) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (u *usecase) DeleteChat(deleter, id int) error {
	if err := u.repo.Delete(deleter, id); err != nil {
		if errors.Is(err, models.ErrPermisionDenied) || errors.Is(err, models.ErrNotFound) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (u *usecase) UpdateChat(chat models.Chatroom) error {
	if err := u.repo.Update(chat); err != nil {
		if errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrNotFound) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (u *usecase) ValidatePassword(id int, password string) (bool, error) {
	pwd, err := u.repo.GetRoomPassword(id)
	if err != nil && errors.Is(err, models.ErrNotFound) {
		return false, err
	} else if err != nil && !errors.Is(err, models.ErrNotFound) {
		log.Fatal(err)
	}

	if !hash.Hshr.Validate(pwd, password) {
		return false, models.ErrPermisionDenied
	}

	return true, nil
}

func (c *usecase) EnterChat(uid, chatroomID int) error {
	if err := c.repo.AddUserToChatroom(uid, chatroomID); err != nil {
		if errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrUserAlreadyInChat) {
			return err
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func (c *usecase) LeaveChat(uid, chatroomID int) error {
	if err := c.repo.RemoveUserFromChatroom(uid, chatroomID); err != nil {
		log.Fatal(err)
	}

	return nil
}
