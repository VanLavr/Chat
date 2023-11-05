package schema

type UserChat struct {
	UserID     int
	ChatroomID int
	ID         int `gorm:"primarykey"`
}
