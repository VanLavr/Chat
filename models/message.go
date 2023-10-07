package models

import "time"

type Message struct {
	ID         int `gorm:"primarykey"`
	UserID     int `gorm:"foreignkey"`
	ChatroomID int `gorm:"foreignkey"`
	Content    string
	Sended     time.Time
}

type MessageRepository interface {
	Fetch(limit int) ([]Message, error)
	FetchOne(id int) (Message, error)
	FetchByUserID(limit, id int) ([]Message, error)
	FetchByChatroomID(limit, id int) ([]Message, error)
	Store(Message *Message) error
	Update(Message *Message) error
	Delete(id int) error
}
