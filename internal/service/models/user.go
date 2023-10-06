package models

type User struct {
	ID         int `gorm:"primarykey"`
	Name       string
	Password   string
	Email      string
	ChatRoomID int `gorm:"foreignkey"`
}

type ChatRoom struct {
	ID   int `gorm:"primarykey"`
	Name string
}

type Message struct {
	ID         int `gorm:"primarykey"`
	UserID     int `gorm:"foreignkey"`
	ChatRoomID int `gorm:"foreignkey"`
	Content    string
}
