package delivery

import (
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	usecase models.MessageUsecase
	jwtmiddleware.JwtMiddleware
}

func NewMessageHandler(e *echo.Echo, u models.MessageUsecase) {
	jwt := jwtmiddleware.NewJwtMiddlware()

	mh := &MessageHandler{usecase: u, JwtMiddleware: *jwt}

	e.GET("/messages/:limit", mh.ValidateToken(mh.GetMessages))
	e.GET("/messages/:user/:limit", mh.ValidateToken(mh.GetUserMessages))
	e.GET("/messages/:chat/:limit", mh.ValidateToken(mh.GetChatMessages))
	e.PUT("/message", mh.ValidateToken(mh.UpdateMessage))
	e.DELETE("/message/:id", mh.ValidateToken(mh.DeleteMessage))
}

func (m *MessageHandler) GetMessages(e echo.Context) error {
	sLimit := e.Param("limit")
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	messages, err := m.usecase.GetMessages(limit)
	if err != nil {
		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Oops, smth went wrong...",
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: messages,
	})
}

func (m *MessageHandler) GetUserMessages(e echo.Context) error {
	sUid := e.Param("user")
	sLimit := e.Param("limit")

	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	uid, err := strconv.Atoi(sUid)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	messages, err := m.usecase.GetUserMessages(limit, uid)
	if err != nil {
		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Oops, smth went wrong",
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: messages,
	})
}

func (m *MessageHandler) GetChatMessages(e echo.Context) error {
	sCid := e.Param("chat")
	sLimit := e.Param("limit")

	limit, err := strconv.Atoi(sLimit)
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

	messages, err := m.usecase.GetChatMessages(limit, cid)
	if err != nil {
		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Oops, smth went wrong",
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: messages,
	})
}

func (m *MessageHandler) UpdateMessage(e echo.Context) error {
	var message models.Message

	err := e.Bind(&message)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err = m.usecase.UpdateMessage(message)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrEmptyFields) {
			return e.JSON(400, models.Response{
				Message: "Failure",
				Content: err,
			})
		}
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "Message updated",
	})
}

func (m *MessageHandler) DeleteMessage(e echo.Context) error {
	sId := e.Param("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err = m.usecase.DeleteMessage(id)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err,
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "Message deleted",
	})
}
