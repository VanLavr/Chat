package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	ur := NewChatroomRepository(schema.NewStorage())

	// test with unexisting Chatroom (empty fields)
	if err := ur.Store(&models.Chatroom{}); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}

	// test with existing Chatroom
	print("CHATROOM ADDED")
	ur.Store(&models.Chatroom{Name: "BOBBY's chat", Password: "123"})
	if err := ur.Store(&models.Chatroom{Name: "BOBBY's chat", Password: "123"}); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	ur := NewChatroomRepository(schema.NewStorage())

	// test with unexisting Chatroom
	if err := ur.Update(&models.Chatroom{ID: 0, Name: "dfks", Password: "jkdf"}); err != models.ErrNotFound {
		t.Fail()
		log.Println(err, "OHO")
	} else {
		log.Println(err.Error())
	}

	// test with unexisting Chatroom (empty fields)
	if err := ur.Update(&models.Chatroom{}); err == nil {
		t.Fail()
		log.Println(err)
	} else {
		log.Println(err.Error())
	}
}

func TestDelete(t *testing.T) {
	ur := NewChatroomRepository(schema.NewStorage())

	// test with unexisting Chatroom
	if err := ur.Delete(0); err != models.ErrNotFound {
		t.Fail()
	} else {
		log.Println(err.Error())
	}
}
