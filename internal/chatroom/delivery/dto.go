package delivery

type EnterChatDTO struct {
	Uid          int
	Cid          int
	RoomPassword string
}

type DeleteChatDTO struct {
	Uid int `json:"uid"`
	Cid int `json:"cid"`
}
