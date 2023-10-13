package delivery

import (
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChatroomHandler struct {
	usecase models.ChatroomUsecase
	jwtmiddleware.JwtMiddleware
}

func NewChatroomHandler(e *echo.Echo, u models.ChatroomUsecase) {
	jwt := jwtmiddleware.NewJwtMiddlware()

	ch := &ChatroomHandler{usecase: u, JwtMiddleware: *jwt}

	e.POST("/user/enterChatroom", ch.ValidateToken(ch.EnterChatroom))
	e.GET("/user/:uid/leaveRoom/:chatroom_id", ch.ValidateToken(ch.LeaveChatroom))
}

func (c *ChatroomHandler) EnterChatroom(e echo.Context) error {
	var UserChat struct {
		Uid          int
		Cid          int
		RoomPassword string
	}

	err := e.Bind(&UserChat)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	valid, err := c.usecase.ValidatePassword(UserChat.Cid, UserChat.RoomPassword)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err,
		})
	}

	if !valid {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrPermisionDenied,
		})
	}

	err = c.usecase.EnterChat(UserChat.Uid, UserChat.Cid)
	if err != nil && (errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrUserAlreadyInChat)) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrNotFound.Error() + " or " + models.ErrUserAlreadyInChat.Error(),
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "user has entered in chatroom",
	})
}

func (c *ChatroomHandler) LeaveChatroom(e echo.Context) error {
	sUID := e.Param("uid")
	sCID := e.Param("chatroom_id")

	uid, err := strconv.Atoi(sUID)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	cid, err := strconv.Atoi(sCID)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err = c.usecase.LeaveChat(uid, cid)
	if err != nil && (errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrUserAlreadyInChat)) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrNotFound.Error() + " or " + models.ErrUserAlreadyInChat.Error(),
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "user has leaved in chatroom",
	})
}
