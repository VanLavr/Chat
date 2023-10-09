package usecase

type CreateChatroomDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	ID       int    `json:"json"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type DeleteUserDTO struct {
	ID      int `json:"id"`
	Deleter int `json:"deleter_id"`
}
