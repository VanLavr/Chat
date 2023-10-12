package delivery

import (
	"chat/models"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase models.UserUsecase
}

func NewUserHandler(e *echo.Echo, u models.UserUsecase) {
	uh := &UserHandler{usecase: u}

	e.GET("/users", uh.GetUsers)
	e.GET("/user/:id", uh.GetUser)
}

func (u *UserHandler) GetUsers(e echo.Context) error {
	limit := e.QueryParam("limit")
	lm, err := strconv.Atoi(limit)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrBadParamInput,
		})
	}

	users := u.usecase.GetUsers(lm)
	return e.JSON(200, models.Response{
		Message: "Success",
		Content: users,
	})
}

func (u *UserHandler) GetUser(e echo.Context) error {
	stringID := e.QueryParam("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrBadParamInput,
		})
	}

	user, err := u.usecase.GetById(id)
	if err != nil && (errors.Is(err, models.ErrNotFound)) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrNotFound,
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: user,
	})
}
