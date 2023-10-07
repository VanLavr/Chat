package models

type ChatRoom struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Password string
}

type ChatroomRepository interface {
	Fetch(limit int) ([]ChatRoom, error)
	FetchOne(id int) (ChatRoom, error)
	Store(*ChatRoom) error
	Update(c *ChatRoom) error
	Delete(id int) error
}
