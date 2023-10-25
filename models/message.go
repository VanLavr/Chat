package models

import "time"

const TimeLayout = "2006-01-02T15:04:05.000Z"

type Message struct {
	ID         int       `gorm:"primarykey" json:"id"`
	UserID     int       `gorm:"foreignkey" json:"user_id"`
	ChatroomID int       `gorm:"foreignkey" json:"chat_id"`
	Content    string    `json:"content"`
	Sended     time.Time `json:"sended"`
}

type MessageRepository interface {
	Fetch(limit int) ([]Message, error)
	FetchOne(id int) (Message, error)
	FetchByUserID(limit, id int) ([]Message, error)
	FetchByChatroomID(limit, id int) ([]Message, error)
	Store(Message Message) error
	StorePhoto(Message Message) (string, error)
	FindPhoto(message Message) (string, error)
	Update(Message Message) error
	Delete(id int) error
	DeletePhoto(id string) (int64, error)
}

type MessageUsecase interface {
	GetChatMessages(limit, id int) ([]Message, error)
	GetUserMessages(limit, id int) ([]Message, error)
	GetMessages(limit int) ([]Message, error)
	GetById(id int) (Message, error)
	CreateMessage(message Message) error
	StorePhoto(message Message) (string, error)
	FindPhoto(message Message) (string, error)
	UpdateMessage(message Message) error
	DeleteMessage(id int) error
	DeletePhoto(id string) (int64, error)
}
