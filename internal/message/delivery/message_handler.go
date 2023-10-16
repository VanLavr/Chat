package delivery

import (
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
	"chat/pkg/logger"
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	usecase models.MessageUsecase
	jwtmiddleware.JwtMiddleware
}

func Register(e *echo.Echo, u models.MessageUsecase) {
	jwt := jwtmiddleware.NewJwtMiddlware()

	mh := &MessageHandler{usecase: u, JwtMiddleware: *jwt}

	e.GET("/messages/:limit", mh.ValidateToken(mh.GetMessages))
	e.GET("/messages/user/:user/:limit", mh.ValidateToken(mh.GetUserMessages))
	e.GET("/messages/chat/:chat/:limit", mh.ValidateToken(mh.GetChatMessages))
	e.PUT("/message", mh.ValidateToken(mh.UpdateMessage))
	e.DELETE("/message/:id", mh.ValidateToken(mh.DeleteMessage))
}

func (m *MessageHandler) GetMessages(e echo.Context) error {
	sLimit := e.Param("limit")
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		logger.FileLogger.Info("/messages/:limit [GET]")
		logger.STDLogger.Info("/messages/:limit [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	messages, err := m.usecase.GetMessages(limit)
	if err != nil {
		logger.FileLogger.Info("/messages/:limit [GET]")
		logger.STDLogger.Info("/messages/:limit [GET]")

		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Oops, smth went wrong...",
		})
	}

	logger.FileLogger.Info("/messages/:limit [GET]")
	logger.STDLogger.Info("/messages/:limit [GET]")

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
		logger.FileLogger.Info("/messages/:user/:limit [GET]")
		logger.STDLogger.Info("/messages/:user/:limit [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	uid, err := strconv.Atoi(sUid)
	if err != nil {
		logger.FileLogger.Info("/messages/:user/:limit [GET]")
		logger.STDLogger.Info("/messages/:user/:limit [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	messages, err := m.usecase.GetUserMessages(limit, uid)
	if err != nil {
		logger.FileLogger.Info("/messages/:user/:limit [GET]")
		logger.STDLogger.Info("/messages/:user/:limit [GET]")

		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Oops, smth went wrong",
		})
	}

	logger.FileLogger.Info("/messages/:user/:limit [GET]")
	logger.STDLogger.Info("/messages/:user/:limit [GET]")

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
		logger.FileLogger.Info("/messages/:chat/:limit [GET]")
		logger.STDLogger.Info("/messages/:chat/:limit [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	cid, err := strconv.Atoi(sCid)
	if err != nil {
		logger.FileLogger.Info("/messages/:chat/:limit [GET]")
		logger.STDLogger.Info("/messages/:chat/:limit [GET]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	messages, err := m.usecase.GetChatMessages(limit, cid)
	if err != nil {
		logger.FileLogger.Info("/messages/:chat/:limit [GET]")
		logger.STDLogger.Info("/messages/:chat/:limit [GET]")

		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Oops, smth went wrong",
		})
	}

	logger.FileLogger.Info("/messages/:chat/:limit [GET]")
	logger.STDLogger.Info("/messages/:chat/:limit [GET]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: messages,
	})
}

func (m *MessageHandler) UpdateMessage(e echo.Context) error {
	var message models.Message

	err := e.Bind(&message)
	if err != nil {
		logger.FileLogger.Info("/message [PUT]")
		logger.STDLogger.Info("/message [PUT]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err = m.usecase.UpdateMessage(message)
	if err != nil {
		logger.FileLogger.Info("/message [PUT]")
		logger.STDLogger.Info("/message [PUT]")

		if errors.Is(err, models.ErrNotFound) || errors.Is(err, models.ErrEmptyFields) {
			return e.JSON(400, models.Response{
				Message: "Failure",
				Content: err.Error(),
			})
		}
	}

	logger.FileLogger.Info("/message [PUT]")
	logger.STDLogger.Info("/message [PUT]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "Message updated",
	})
}

func (m *MessageHandler) DeleteMessage(e echo.Context) error {
	sId := e.Param("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		logger.FileLogger.Info("/message [DELETE]")
		logger.STDLogger.Info("/message [DELETE]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	err = m.usecase.DeleteMessage(id)
	if err != nil {
		logger.FileLogger.Info("/message [DELETE]")
		logger.STDLogger.Info("/message [DELETE]")

		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	logger.FileLogger.Info("/message [DELETE]")
	logger.STDLogger.Info("/message [DELETE]")

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: "Message deleted",
	})
}
