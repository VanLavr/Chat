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
	Fetch(int) ([]Message, error)
	FetchOne(int) (Message, error)
	FetchByUserID(int, int) ([]Message, error)
	FetchByChatroomID(int, int) ([]Message, error)
	Store(Message) error
	StorePhoto(Message) (string, error)
	FindPhoto(Message) (string, error)
	Update(Message) error
	Delete(int) error
	DeletePhoto(string) (int64, error)
}

type MessageUsecase interface {
	GetChatMessages(int, int) ([]Message, error)
	GetUserMessages(int, int) ([]Message, error)
	GetMessages(int) ([]Message, error)
	GetById(int) (Message, error)
	CreateMessage(Message) error
	StorePhoto(Message) (string, error)
	FindPhoto(Message) (string, error)
	UpdateMessage(Message) error
	DeleteMessage(int) error
	DeletePhoto(string) (int64, error)
}
