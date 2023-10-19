package schema

type UserChat struct {
	UserID     int
	ChatroomID int
	ID         int `gorm:"primarykey"`
}

type Param struct {
	UserID           int
	ChatroomID       int
	ChatroomPassword string
}

type DeleteChat struct {
	DeleterID int
	Chatroom  int
}
