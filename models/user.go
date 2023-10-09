package models

type User struct {
	ID         int `gorm:"primarykey"`
	Name       string
	Password   string
	IsAdmin    bool `gorm:"default:false"`
	RoomsOwned int
}

type UserRepository interface {
	Fetch(limit int) ([]User, error)
	FetchOne(id int) (User, error)
	FetchFewCertain(id ...int) ([]User, error)
	AddUserToChatroom(uid, chatId int) error
	RemoveUserFromChatroom(uid, chatId int) error
	Store(user *User) error
	Update(user *User) error
	Delete(id int) error
}
