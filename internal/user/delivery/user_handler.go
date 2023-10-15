package delivery

import (
	"chat/internal/message/repository/postgres"
	"chat/internal/message/usecase"
	schema "chat/migrations"
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
	"chat/pkg/logger"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase models.UserUsecase
	hub     []models.User
	jwtmiddleware.JwtMiddleware
	websocket.Upgrader
}

func Register(e *echo.Echo, u models.UserUsecase) {
	jwt := jwtmiddleware.NewJwtMiddlware()
	upd := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	uh := &UserHandler{usecase: u, JwtMiddleware: *jwt, Upgrader: upd}

	e.GET("/users/:limit", uh.GetUsers)
	e.GET("/user/:id", uh.GetUser)
	e.POST("/user", uh.CreateUser)
	e.PUT("/user", uh.ValidateToken(uh.UpdateUser))
	e.DELETE("/user", uh.ValidateToken(uh.DeleteUser))
	e.POST("/user/jwt", uh.GetJWT)
	e.GET("/ws/start/:uid/:cid", uh.ValidateToken(uh.Join))
}

func (u *UserHandler) GetUsers(e echo.Context) error {
	limit := e.Param("limit")
	lm, err := strconv.Atoi(limit)
	if err != nil {
		logger.FileLogger.Info("/users/:limit [GET]")
		logger.STDLogger.Info("/users/:limit [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrBadParamInput.Error(),
		})
	}

	users := u.usecase.GetUsers(lm)

	logger.FileLogger.Info("/users/:limit [GET]")
	logger.STDLogger.Info("/users/:limit [GET]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: users,
	})
}

func (u *UserHandler) GetUser(e echo.Context) error {
	stringID := e.Param("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		logger.FileLogger.Info("/user/:id [GET]")
		logger.STDLogger.Info("/user/:id [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrBadParamInput.Error(),
		})
	}

	user, err := u.usecase.GetById(id)
	if err != nil && (errors.Is(err, models.ErrNotFound)) {
		logger.FileLogger.Info("/user/:id [GET]")
		logger.STDLogger.Info("/user/:id [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrNotFound.Error(),
		})
	}

	logger.FileLogger.Info("/user/:id [GET]")
	logger.STDLogger.Info("/user/:id [GET]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: user,
	})
}

func (u *UserHandler) CreateUser(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		logger.FileLogger.Info("/user [POST]")
		logger.STDLogger.Info("/user [POST]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	if user.Name == "" || user.Password == "" {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrEmptyFields.Error(),
		})
	}

	err := u.usecase.CreateUser(user)
	if err != nil && (errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrAlreadyExists)) {
		logger.FileLogger.Info("/user [POST]")
		logger.STDLogger.Info("/user [POST]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	} else {

		logger.FileLogger.Info("/user [POST]")
		logger.STDLogger.Info("/user [POST]")

		return e.JSON(200, models.Response{
			Message: "Success",
			Content: "user created",
		})
	}
}

func (u *UserHandler) UpdateUser(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		logger.FileLogger.Info("/user [PUT]")
		logger.STDLogger.Info("/user [PUT]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	if user.Password == "" {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrPermisionDenied.Error(),
		})
	}

	ok, err := u.usecase.ValidatePassword(user.ID, user.Password)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	if !ok {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrPermisionDenied.Error(),
		})
	}

	err = u.usecase.UpdateUser(user)
	if err != nil && (errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrNotFound)) {
		logger.FileLogger.Info("/user [PUT]")
		logger.STDLogger.Info("/user [PUT]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	} else {

		logger.FileLogger.Info("/user [PUT]")
		logger.STDLogger.Info("/user [PUT]")

		return e.JSON(200, models.Response{
			Message: "Success",
			Content: "user updated",
		})
	}
}

func (u *UserHandler) DeleteUser(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		logger.FileLogger.Info("/user [DELETE]")
		logger.STDLogger.Info("/user [DELETE]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	ok, err := u.usecase.ValidatePassword(user.ID, user.Password)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	if !ok {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrPermisionDenied.Error(),
		})
	}

	err = u.usecase.DeleteUser(user.ID)
	if err != nil && (errors.Is(err, models.ErrNotFound)) {
		logger.FileLogger.Info("/user [DELETE]")
		logger.STDLogger.Info("/user [DELETE]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrNotFound.Error(),
		})
	} else {

		logger.FileLogger.Info("/user [DELETE]")
		logger.STDLogger.Info("/user [DELETE]")

		return e.JSON(200, models.Response{
			Message: "Success",
			Content: "user deleted",
		})
	}
}

func (u *UserHandler) GetJWT(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		logger.FileLogger.Info("/user/jwt [GET]")
		logger.STDLogger.Info("/user/jwt [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	ok, err := u.usecase.ValidatePassword(user.ID, user.Password)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	if !ok {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrPermisionDenied.Error(),
		})
	}

	logger.FileLogger.Info("/user/jwt [GET]")
	logger.STDLogger.Info("/user/jwt [GET]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: u.GenerateToken(user),
	})
}

func (u *UserHandler) Join(e echo.Context) error {
	sUid := e.Param("uid")
	sCid := e.Param("cid")

	uid, err := strconv.Atoi(sUid)
	if err != nil {
		logger.STDLogger.Info("/websocket/start")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	cid, err := strconv.Atoi(sCid)
	if err != nil {
		logger.STDLogger.Info("/websocket/start")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	if !u.usecase.ValidateIncommer(uid, cid) {
		logger.STDLogger.Info("/websocket/start")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrPermisionDenied.Error(),
		})
	}

	user, err := u.usecase.GetById(uid)
	if err != nil {
		logger.STDLogger.Info("/websocket/start")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	user.CurrentChatroomID = cid

	conn, err := u.Upgrade(e.Response().Writer, e.Request(), nil)
	if err != nil {
		logger.STDLogger.Fatal(err.Error())
	}

	user.Connection = conn

	u.hub = append(u.hub, user)

	// go u.readMessage(&user)
	go func(user *models.User) {
		if user.ID == 0 {
			logger.STDLogger.Fatal(errors.New("can not resolve user").Error())
		}

		u.readMessage(user)
	}(&user)

	logger.STDLogger.Info("/websocket/start")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "connection established",
	})

}

func (u *UserHandler) readMessage(user *models.User) {
	for {
		msgT, message, err := user.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway, websocket.CloseProtocolError, websocket.CloseNoStatusReceived) {
				log.Println("abnormal disconnect...")
				break
			} else {
				log.Println(err)
				break
			}
		}

		u.multicast(msgT, message, user)
	}
}

func (u *UserHandler) multicast(msgType int, message []byte, user *models.User) {
	Mrepo := postgres.NewMessageRepository(schema.NewStorage())
	Musecase := usecase.NewUsecase(Mrepo)

	err := Musecase.CreateMessage(models.Message{
		UserID:     user.ID,
		ChatroomID: user.CurrentChatroomID,
		Content:    string(message),
	})

	logger.STDLogger.Info(string(message))

	if err != nil {
		return
	}

	for i := 0; i < len(u.hub); i++ {
		if u.hub[i].CurrentChatroomID == user.CurrentChatroomID && u.hub[i].ID != user.ID {
			if err := u.hub[i].Connection.WriteMessage(msgType, message); err != nil {
				logger.STDLogger.Fatal(err.Error())
			}
		}
	}
}
