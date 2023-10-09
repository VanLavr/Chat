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
	if err := ur.Store(&models.User{}); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}

	// test with existing user
	ur.Store(&models.User{Name: "BOBBY", Password: "123"})
	if err := ur.Store(&models.User{Name: "BOBBY", Password: "123"}); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	ur := NewUserRepository(schema.NewStorage())

	// test with unexisting user
	if err := ur.Update(&models.User{ID: 0, Name: "dfks", Password: "jkdf"}); err != models.ErrNotFound {
		t.Fail()
		log.Println(err, "OHO")
	} else {
		log.Println(err.Error())
	}

	// test with unexisting user (empty fields)
	if err := ur.Update(&models.User{}); err == nil {
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
}

func TestAddToChatroom(t *testing.T) {
	ur := NewUserRepository(schema.NewStorage())

	// test with unexisting user
	if err := ur.AddUserToChatroom(2, 2); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}
}
