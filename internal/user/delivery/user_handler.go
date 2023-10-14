package delivery

import (
	"chat/internal/message/repository/postgres"
	"chat/internal/message/usecase"
	schema "chat/migrations"
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
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

func NewUserHandler(e *echo.Echo, u models.UserUsecase) {
	jwt := jwtmiddleware.NewJwtMiddlware()
	upd := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	uh := &UserHandler{usecase: u, JwtMiddleware: *jwt, Upgrader: upd}

	e.GET("/users", uh.GetUsers)
	e.GET("/user/:id", uh.GetUser)
	e.POST("/user", uh.CreateUser)
	e.PUT("/user", uh.ValidateToken(uh.UpdateUser))
	e.DELETE("/user", uh.ValidateToken(uh.DeleteUser))
	e.GET("/user/jwt", uh.GetJWT)
	e.GET("/ws/start/:uid/:cid", uh.ValidateToken(uh.Join))
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
			Content: err,
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

func (u *UserHandler) Join(e echo.Context) error {
	sUid := e.Param("uid")
	sCid := e.Param("cid")

	uid, err := strconv.Atoi(sUid)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	cid, err := strconv.Atoi(sCid)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	if !u.usecase.ValidateIncommer(uid, cid) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrPermisionDenied,
		})
	}

	user, err := u.usecase.GetById(uid)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err,
		})
	}

	user.CurrentChatroomID = cid

	conn, err := u.Upgrade(e.Response().Writer, e.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}

	user.Connection = conn

	u.hub = append(u.hub, user)

	// go u.readMessage(&user)
	go func(user *models.User) {
		if user.ID == 0 {
			log.Fatal(errors.New("can not resolve user"))
		}

		u.readMessage(user)
	}(&user)

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

		log.Println(string(message))

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

	if err != nil {
		return
	}

	for i := 0; i < len(u.hub); i++ {
		if u.hub[i].CurrentChatroomID == user.CurrentChatroomID && u.hub[i].ID != user.ID {
			if err := u.hub[i].Connection.WriteMessage(msgType, message); err != nil {
				log.Fatal(err)
			}
		}
	}
}
