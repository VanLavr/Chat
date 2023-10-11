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
	Store(chat Chatroom) error
	Update(chat Chatroom) error
	Delete(deleter, id int) error
}

type ChatroomUsecase interface {
	GetById(id int) (Chatroom, error)
	Get(limit int) ([]Chatroom, error)
	CreateChat(chatroom Chatroom) error
	DeleteChat(deleter, id int) error
	UpdateChat(chat Chatroom) error
}
