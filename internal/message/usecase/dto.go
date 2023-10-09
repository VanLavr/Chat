package usecase

type CreateMessageDTO struct {
	UserID     int    `json:"user_id"`
	ChatroomID int    `json:"chatroom_id"`
	Content    string `json:"content"`
}

type UpdateMessageDTO struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type DeleteMessageDTO struct {
	ID int `json:"id"`
}
