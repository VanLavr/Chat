package models

type Chatroom struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Password string
}

type ChatroomRepository interface {
	Fetch(limit int) ([]Chatroom, error)
	FetchOne(id int) (Chatroom, error)
	Store(*Chatroom) error
	Update(c *Chatroom) error
	Delete(id int) error
}
