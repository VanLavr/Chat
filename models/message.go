package models

import "time"

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
	Update(Message Message) error
	Delete(id int) error
}

type MessageUsecase interface {
	GetChatMessages(limit, id int) ([]Message, error)
	GetUserMessages(limit, id int) ([]Message, error)
	GetMessages(limit int) ([]Message, error)
	GetById(id int) (Message, error)
	CreateMessage(message Message) error
	UpdateMessage(message Message) error
	DeleteMessage(id int) error
}
