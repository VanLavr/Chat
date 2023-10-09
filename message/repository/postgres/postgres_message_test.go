package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	ur := NewMessageRepository(schema.NewStorage())

	// test with unexisting Message (empty fields)
	if err := ur.Store(&models.Message{}); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}

	// test with existing Message
	print("Message ADDED")
	ur.Store(&models.Message{Content: "test", UserID: 11, ChatroomID: 2})
	if err := ur.Store(&models.Message{Content: "test", UserID: 11, ChatroomID: 2}); err == nil {
		t.Fail()
	} else {
		log.Println(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	ur := NewMessageRepository(schema.NewStorage())

	// test with unexisting Message
	if err := ur.Update(&models.Message{ID: 0, Content: "hdjsgfdjtest", UserID: 11, ChatroomID: 2}); err != models.ErrNotFound {
		t.Fail()
		log.Println(err, "OHO")
	} else {
		log.Println(err.Error())
	}

	// test with unexisting Message (empty fields)
	if err := ur.Update(&models.Message{}); err == nil {
		t.Fail()
		log.Println(err)
	} else {
		log.Println(err.Error())
	}
}

func TestDelete(t *testing.T) {
	ur := NewMessageRepository(schema.NewStorage())

	// test with unexisting Message
	if err := ur.Delete(0); err != models.ErrNotFound {
		t.Fail()
	} else {
		log.Println(err.Error())
	}
}
