package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	ur := NewUserRepository(schema.NewStorage())

	// test with unexisting user (empty fields)
	if err := ur.Store(models.User{}); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}

	// test with existing user
	if err := ur.Store(models.User{Name: "BOBBY", Password: "123"}); err != nil {
		log.Println(err, "HUI")
	}
	if err := ur.Store(models.User{Name: "BOBBY", Password: "123"}); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	ur := NewUserRepository(schema.NewStorage())

	// test with unexisting user
	if err := ur.Update(models.User{ID: 0, Name: "dfks", Password: "jkdf"}); err != models.ErrNotFound {
		t.Fail()
		log.Println(err)
	} else {
		log.Println(err.Error())
	}

	// test with existing user (empty fields)
	if err := ur.Update(models.User{ID: 1}); err == nil {
		t.Fail()
		log.Println(err)
	} else {
		log.Println(err.Error())
	}
}

func TestDelete(t *testing.T) {
	ur := NewUserRepository(schema.NewStorage())

	// test with unexisting user
	if err := ur.Delete(0); err != models.ErrNotFound {
		t.Fail()
	} else {
		log.Println(err.Error())
	}

	if err := ur.Delete(1); err != nil {
		log.Println(err)
		t.Fail()
	}
}

func TestAddToChatroom(t *testing.T) {
	ur := NewUserRepository(schema.NewStorage())

	// test with unexisting user and existing chatroom
	if err := ur.AddUserToChatroom(0, 1); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}

	// test with exsisting user and unexisting chatroom
	if err := ur.AddUserToChatroom(2, 0); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}

	if err := ur.AddUserToChatroom(2, 1); err != nil {
		log.Println(err, "jopa")
		t.Fail()
	}
}
