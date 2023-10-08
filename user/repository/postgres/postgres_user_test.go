package postgres

import (
	schema "chat/migrations"
	"chat/models"
	"testing"
)

func TestUpdate(t *testing.T) {
	ur := NewUserRepository(schema.NewStorage())

	// test with unexistinguser
	if err := ur.Update(&models.User{}); err != models.ErrNotFound {
		t.Fail()
	}
}
