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
	Store(user User) error
	Update(user User) error
	Delete(id int) error
	GetChatters() []User
	GetUserPassword(id int) (string, error)
	BeforeJoin(uid, cid int) bool
}

type UserUsecase interface {
	GetById(uid int) (User, error)
	GetUsers(limit int) []User
	CreateUser(user User) error
	UpdateUser(user User) error
	DeleteUser(id int) error
	MakeHub() []User
	ValidatePassword(uid int, password string) (bool, error)
	ValidateIncommer(uid, cid int) bool
}
