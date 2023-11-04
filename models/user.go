package models

import (
	"github.com/gorilla/websocket"
)

type User struct {
	ID                int             `gorm:"primarykey" json:"id"`
	Name              string          `json:"name"`
	Password          string          `json:"password"`
	IsAdmin           bool            `gorm:"default:false" json:"admin"`
	RoomsOwned        int             `json:"rooms_owned"`
	CurrentChatroomID int             `gorm:"-"`
	Connection        *websocket.Conn `gorm:"-" json:"-"`
}

type UserRepository interface {
	Fetch(int) ([]User, error)
	FetchOne(int) (User, error)
	FetchFewCertain(...int) ([]User, error)
	Store(User) error
	Update(User) error
	Delete(int) error
	GetChatters() []User
	GetUserPassword(int) (string, error)
	GetuserName(int) (string, error)
	BeforeJoin(int, int) bool
}

type UserUsecase interface {
	GetById(int) (User, error)
	GetUsers(int) []User
	CreateUser(User) error
	UpdateUser(User) error
	DeleteUser(int) error
	MakeHub() []User
	ValidatePassword(int, string) (bool, error)
	ValidateUsername(int, string) (bool, error)
	ValidateIncommer(int, int) bool
}
