package delivery

import (
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase models.UserUsecase
	jwtmiddleware.JwtMiddleware
}

func NewUserHandler(e *echo.Echo, u models.UserUsecase) {
	jwt := jwtmiddleware.NewJwtMiddlware()
	uh := &UserHandler{usecase: u, JwtMiddleware: *jwt}

	e.GET("/users", uh.GetUsers)
	e.GET("/user/:id", uh.GetUser)
	e.POST("/user", uh.CreateUser)
	e.PUT("/user", uh.ValidateToken(uh.UpdateUser))
	e.DELETE("/user", uh.ValidateToken(uh.DeleteUser))
	e.GET("user/jwt", uh.GetJWT)
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

func (u *UserHandler) CreateUser(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err := u.usecase.CreateUser(user)
	if err != nil && (errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrAlreadyExists)) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrAlreadyExists.Error() + "or" + models.ErrEmptyFields.Error(),
		})
	} else {
		return e.JSON(200, models.Response{
			Message: "Success",
			Content: "user created",
		})
	}
}

func (u *UserHandler) UpdateUser(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err := u.usecase.UpdateUser(user)
	if err != nil && (errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrNotFound)) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrNotFound.Error() + "or" + models.ErrEmptyFields.Error(),
		})
	} else {
		return e.JSON(200, models.Response{
			Message: "Success",
			Content: "user updated",
		})
	}
}

func (u *UserHandler) DeleteUser(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err := u.usecase.DeleteUser(user.ID)
	if err != nil && (errors.Is(err, models.ErrNotFound)) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrNotFound.Error(),
		})
	} else {
		return e.JSON(200, models.Response{
			Message: "Success",
			Content: "user deleted",
		})
	}
}

func (u *UserHandler) GetJWT(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: u.GenerateToken(user),
	})
}
