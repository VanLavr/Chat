package models

type User struct {
	ID       int    `gorm:"primarykey"`
	Name     string `gorm:"uniqueIndex:udx_name"`
	Password string
}

func (u User) Create(Name, Password string) {

}

func (u User) Update(Name, Password string, id int) {

}

func (u User) Delete(id int) {

}

type ChatRoom struct {
	ID       int    `gorm:"primarykey"`
	Name     string `gorm:"uniqueIndex:cdx_name"`
	Password string
}

func (c ChatRoom) Create(Name, Password string) {

}

func (c ChatRoom) Update(Name, Password string, id int) {

}

func (c ChatRoom) Delete(id int) {

}

type Message struct {
	ID         int `gorm:"primarykey"`
	UserChatID int `gorm:"foreignkey"`
	Content    string
}

func (m Message) Create(Content string, UserChatID int) {

}

func (m Message) Update(Content string, id int) {

}

func (m Message) Delete(id int) {

}

type UserChat struct {
	ID         int `gorm:"primarykey"`
	UserID     int `gorm:"foreignkey"`
	ChatRoomID int `gorm:"foreignkey"`
}

func (uc UserChat) Create(UserID, ChatRoomID int) {

}

func (uc UserChat) Delete(id int) {

}
