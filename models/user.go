package models

type User struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Password string
}

type UserRepository interface {
	Fetch(limit int) ([]User, error)
	FetchOne(id int) (User, error)
	FetchFewCertain(id ...int) ([]User, error)
	Store(*User) error
	Update(m *User) error
	Delete(id int) error
}
