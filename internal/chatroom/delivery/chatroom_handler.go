package delivery

import (
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
	"chat/pkg/logger"
	"errors"
	"fmt"
	"strconv"

	_ "chat/docs"

	"github.com/labstack/echo/v4"
)

type ChatroomHandler struct {
	usecase models.ChatroomUsecase
	jwtmiddleware.JwtMiddleware
}

func Register(e *echo.Echo, u models.ChatroomUsecase) {
	jwt := jwtmiddleware.NewJwtMiddlware()

	ch := &ChatroomHandler{usecase: u, JwtMiddleware: *jwt}

	e.GET("/chatrooms/:limit", ch.GetRooms)
	e.GET("/chatroom/:id", ch.GetById)
	e.POST("/user/enterChatroom", ch.ValidateToken(ch.EnterChatroom))
	e.GET("/user/:uid/leaveRoom/:chatroom_id", ch.ValidateToken(ch.LeaveChatroom))
	e.POST("/chatroom", ch.ValidateToken(ch.CreateChat))
	e.PUT("/chatroom", ch.ValidateToken(ch.UpdateChat))
	e.DELETE("/chatroom", ch.ValidateToken(ch.DeleteChat))
}

//	@Summary		Get chatroom by ID
//	@Description	Get a chatroom by its ID
//	@ID				get-chatroom-by-id
//	@Produce		json
//	@Param			id	path		int	true	"Chatroom ID"
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Router			/chatroom/{id} [get]
func (c *ChatroomHandler) GetById(e echo.Context) error {
	sid := e.Param("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		logger.STDLogger.Info("/chatroom/:id")
		logger.FileLogger.Info("/chatroom/:id")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	chat, err := c.usecase.GetById(id)
	if err != nil {
		logger.STDLogger.Info("/chatroom/:id")
		logger.FileLogger.Info("/chatroom/:id")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	logger.STDLogger.Info("/chatroom/:id")
	logger.FileLogger.Info("/chatroom/:id")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: chat,
	})
}

//	@Summary		Get chatrooms
//	@Description	Get a list of chatrooms
//	@ID				get-chatrooms
//	@Produce		json
//	@Param			limit	path		int	true	"Number of chatrooms to retrieve"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Response
//	@Router			/chatrooms/{limit} [get]
func (c *ChatroomHandler) GetRooms(e echo.Context) error {
	sLimit := e.Param("limit")
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		logger.FileLogger.Info("/chatrooms [GET]")
		logger.STDLogger.Info("/chatrooms [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	chats, err := c.usecase.Get(limit)
	if err != nil {
		logger.FileLogger.Info("/chatrooms [GET]")
		logger.STDLogger.Info("/chatrooms [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	logger.FileLogger.Info("/chatrooms [GET]")
	logger.STDLogger.Info("/chatrooms [GET]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: chats,
	})
}

func (c *ChatroomHandler) EnterChatroom(e echo.Context) error {
	var UserChat struct {
		Uid          int
		Cid          int
		RoomPassword string
	}

	err := e.Bind(&UserChat)
	if err != nil {
		logger.FileLogger.Info("/user/enterChatroom [POST]")
		logger.STDLogger.Info("/user/enterChatroom [POST]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	valid, err := c.usecase.ValidatePassword(UserChat.Cid, UserChat.RoomPassword)
	if err != nil {
		logger.FileLogger.Info("/user/enterChatroom [POST]")
		logger.STDLogger.Info("/user/enterChatroom [POST]")

		fmt.Println(err.Error())
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	if !valid {
		logger.FileLogger.Info("/user/enterChatroom [POST]")
		logger.STDLogger.Info("/user/enterChatroom [POST]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrPermisionDenied.Error(),
		})
	}

	err = c.usecase.EnterChat(UserChat.Uid, UserChat.Cid)
	if err != nil && (errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrUserAlreadyInChat)) {
		logger.FileLogger.Info("/user/enterChatroom [POST]")
		logger.STDLogger.Info("/user/enterChatroom [POST]")

		fmt.Println(err.Error())
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	logger.FileLogger.Info("/user/enterChatroom [POST]")
	logger.STDLogger.Info("/user/enterChatroom [POST]")

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
		logger.FileLogger.Info("/user/:uid/leaveRoom/:chatroom_id [GET]")
		logger.STDLogger.Info("/user/:uid/leaveRoom/:chatroom_id [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	cid, err := strconv.Atoi(sCID)
	if err != nil {
		logger.FileLogger.Info("/user/:uid/leaveRoom/:chatroom_id [GET]")
		logger.STDLogger.Info("/user/:uid/leaveRoom/:chatroom_id [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err = c.usecase.LeaveChat(uid, cid)
	if err != nil && (errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrUserAlreadyInChat)) {
		logger.FileLogger.Info("/user/:uid/leaveRoom/:chatroom_id [GET]")
		logger.STDLogger.Info("/user/:uid/leaveRoom/:chatroom_id [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	logger.FileLogger.Info("/user/:uid/leaveRoom/:chatroom_id [GET]")
	logger.STDLogger.Info("/user/:uid/leaveRoom/:chatroom_id [GET]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "user has leaved in chatroom",
	})
}

func (c *ChatroomHandler) CreateChat(e echo.Context) error {
	var chat models.Chatroom

	err := e.Bind(&chat)
	if err != nil {
		logger.FileLogger.Info("/chatroom [POST]")
		logger.STDLogger.Info("/chatroom [POST]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	if err = c.usecase.CreateChat(chat); err != nil {
		if errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrAlreadyExists) {
			logger.FileLogger.Info("/chatroom [POST]")
			logger.STDLogger.Info("/chatroom [POST]")

			return e.JSON(400, models.Response{
				Message: "Failure",
				Content: err.Error(),
			})
		}
	}

	logger.FileLogger.Info("/chatroom [POST]")
	logger.STDLogger.Info("/chatroom [POST]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "Chatroom created",
	})
}

func (c *ChatroomHandler) UpdateChat(e echo.Context) error {
	var chat models.Chatroom

	err := e.Bind(&chat)
	if err != nil {
		logger.FileLogger.Info("/chatroom [PUT]")
		logger.STDLogger.Info("/chatroom [PUT]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	ok, err := c.usecase.ValidatePassword(chat.ID, chat.Password)
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

	if err = c.usecase.UpdateChat(chat); err != nil {
		if errors.Is(err, models.ErrEmptyFields) || errors.Is(err, models.ErrNotFound) {
			logger.FileLogger.Info("/chatroom [PUT]")
			logger.STDLogger.Info("/chatroom [PUT]")

			return e.JSON(400, models.Response{
				Message: "Failure",
				Content: err.Error(),
			})
		}
	}

	logger.FileLogger.Info("/chatroom [PUT]")
	logger.STDLogger.Info("/chatroom [PUT]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "Chatroom updated",
	})
}

func (c *ChatroomHandler) DeleteChat(e echo.Context) error {
	var deletion struct {
		Uid int `json:"uid"`
		Cid int `json:"cid"`
	}

	err := e.Bind(&deletion)
	if err != nil {
		logger.FileLogger.Info("/chatroom [DELETE]")
		logger.STDLogger.Info("/chatroom [DELETE]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	fmt.Println(deletion)

	err = c.usecase.DeleteChat(deletion.Uid, deletion.Cid)
	if err != nil {
		logger.FileLogger.Info("/chatroom [DELETE]")
		logger.STDLogger.Info("/chatroom [DELETE]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	logger.FileLogger.Info("/chatroom [DELETE]")
	logger.STDLogger.Info("/chatroom [DELETE]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "Chatroom deleted",
	})
}
