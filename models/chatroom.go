package models

type Chatroom struct {
	ID        int    `gorm:"primarykey" json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatorID int    `json:"owner"`
}

type ChatroomRepository interface {
	Fetch(limit int) ([]Chatroom, error)
	FetchOne(id int) (Chatroom, error)
	Store(*Chatroom) error
	Update(c *Chatroom) error
	Delete(id int) error
}

type ChatroomUsecase interface {
	GetById(id int) (Chatroom, error)
	Get(limit int) ([]Chatroom, error)
	CreateChat(chatroom Chatroom) error
	DeleteChat(chat Chatroom) error
	UpdateChat(chat Chatroom) error
}
