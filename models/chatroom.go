package models

type Chatroom struct {
	ID        int    `gorm:"primarykey" json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatorID int    `json:"owner"`
}

type ChatroomRepository interface {
	Fetch(int) ([]Chatroom, error)
	FetchOne(int) (Chatroom, error)
	Store(Chatroom) error
	Update(Chatroom) error
	Delete(int, int) error
	GetRoomPassword(int) (string, error)
	AddUserToChatroom(int, int) error
	RemoveUserFromChatroom(int, int) error
}

type ChatroomUsecase interface {
	GetById(int) (Chatroom, error)
	Get(int) ([]Chatroom, error)
	CreateChat(Chatroom) error
	DeleteChat(int, int) error
	UpdateChat(Chatroom) error
	ValidatePassword(int, string) (bool, error)
	EnterChat(int, int) error
	LeaveChat(int, int) error
}
