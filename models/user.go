package models

import (
	"github.com/gorilla/websocket"
)

type User struct {
	ID                int    `gorm:"primarykey" json:"id"`
	Name              string `json:"name"`
	Password          string `json:"password"`
	IsAdmin           bool   `gorm:"default:false" json:"admin"`
	RoomsOwned        int    `json:"rooms_owned"`
	CurrentChatroomID int
	Connection        *websocket.Conn
}

type UserRepository interface {
	Fetch(limit int) ([]User, error)
	FetchOne(id int) (User, error)
	FetchFewCertain(id ...int) ([]User, error)
	AddUserToChatroom(uid, chatId int) error
	RemoveUserFromChatroom(uid, chatId int) error
	Store(user User) error
	Update(user User) error
	Delete(id int) error
	GetChatters() []User
	GetUserPassword(id int) (string, error)
}

type UserUsecase interface {
	GetById(uid int) (User, error)
	GetUsers(limit int) []User
	EnterChat(uid, chatroomID int) error
	LeaveChat(uid, chatroomID int) error
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(id int) error
	MakeHub() []User
	ValidatePassword(uid int, password string) (bool, error)
}
