package schema

type UserChat struct {
	UserID     int `gorm:"foreignkey"`
	ChatroomID int `gorm:"foreignkey"`
	ID         int `gorm:"primarykey"`
}
