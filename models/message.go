package models

type Message struct {
	ID         int `gorm:"primarykey"`
	UserChatID int `gorm:"foreignkey"`
	Content    string
}

type MessageRepository interface {
	Fetch(limit int) ([]Message, error)
	FetchOne(id int) (Message, error)
	FetchByUserID(limit, id int) (Message, error)
	FetchByChatroomID(limit, id int) (Message, error)
	Store(*Message) error
	Update(m *Message) error
	Delete(id int) error
}
