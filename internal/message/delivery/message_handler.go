package delivery

import (
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
	"chat/pkg/logger"
	"errors"
	"strconv"

	_ "chat/docs"

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

//	@Summary		Get messages
//	@Tags			messages
//	@Description	Retrieve messages with a specified limit
//	@ID				get-messages
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int	true	"Limit of messages to retrieve"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Response
//	@Failure		500		{object}	models.Response
//	@Router			/messages/{limit} [get]
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

//	@Summary		Get user messages
//	@Tags			messages
//	@Description	Retrieve messages for a specific user with a specified limit
//	@ID				get-user-messages
//	@Accept			json
//	@Produce		json
//	@Param			user	path		int	true	"User ID"
//	@Param			limit	path		int	true	"Limit of messages to retrieve"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Response
//	@Failure		500		{object}	models.Response
//	@Router			/messages/{user}/{limit} [get]
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

//	@Summary		Get chat messages
//	@Tags			messages
//	@Description	Retrieve messages for a specific chat with a specified limit
//	@ID				get-chat-messages
//	@Accept			json
//	@Produce		json
//	@Param			chat	path		int	true	"Chat ID"
//	@Param			limit	path		int	true	"Limit of messages to retrieve"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Response
//	@Failure		500		{object}	models.Response
//	@Router			/messages/{chat}/{limit} [get]
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

//	@Summary		Update a message
//	@Tags			messages
//	@Description	Update a message with new content
//	@ID				update-message
//	@Accept			json
//	@Produce		json
//	@Param			message	body		models.Message	true	"Message object"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Response
//	@Router			/message [put]
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

//	@Summary		Delete a message
//	@Tags			messages
//	@Description	Delete a message with a specified ID
//	@ID				delete-message
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Message ID"
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Router			/message/{id} [delete]
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
