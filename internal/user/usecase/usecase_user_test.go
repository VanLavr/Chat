package usecase

import (
	"chat/models"
	mock_models "chat/models/mock"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSave(t *testing.T) {
	Controller := gomock.NewController(t)
	repo := mock_models.NewMockUserRepository(Controller)

	usecase := NewUsecase(repo)

	DataSet := []models.User{
		{
			Name:     "hah",
			Password: "alkfj",
		},
		{
			Name: "fjslkd",
		},
		{
			Password: "a",
		},
		{},
	}

	for _, user := range DataSet {
		err := usecase.CreateUser(user)
		if err != nil && !errors.Is(err, models.ErrEmptyFields) {
			t.Errorf("FAILED error expected: %v. Got: %v", models.ErrEmptyFields, err)
		}
	}

	t.Log("PASSED user creation")
}
