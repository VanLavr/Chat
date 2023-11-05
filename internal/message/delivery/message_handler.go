package delivery

import (
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
	"chat/pkg/logger"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	_ "chat/docs"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
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
	e.POST("/message/upload-photo", mh.ValidateToken(mh.UploadPhoto))
	e.POST("/message/find-photo", mh.ValidateToken(mh.FindPhoto))
	e.DELETE("/message/delete-photo/:id", mh.ValidateToken(mh.DeletePhoto))
}

// @Summary		Get messages
// @Tags			messages
// @Description	Retrieve messages with a specified limit
// @ID				get-messages
// @Accept			json
// @Produce		json
// @Param			limit	path		int	true	"Limit of messages to retrieve"
// @Success		200		{object}	models.Response
// @Failure		400		{object}	models.Response
// @Failure		500		{object}	models.Response
// @Router			/messages/{limit} [get]
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

// @Summary		Get user messages
// @Tags			messages
// @Description	Retrieve messages for a specific user with a specified limit
// @ID				get-user-messages
// @Accept			json
// @Produce		json
// @Param			user	path		int	true	"User ID"
// @Param			limit	path		int	true	"Limit of messages to retrieve"
// @Success		200		{object}	models.Response
// @Failure		400		{object}	models.Response
// @Failure		500		{object}	models.Response
// @Router			/messages/user/{user}/{limit} [get]
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

// @Summary		Get chat messages
// @Tags			messages
// @Description	Retrieve messages for a specific chat with a specified limit
// @ID				get-chat-messages
// @Accept			json
// @Produce		json
// @Param			chat	path		int	true	"Chat ID"
// @Param			limit	path		int	true	"Limit of messages to retrieve"
// @Success		200		{object}	models.Response
// @Failure		400		{object}	models.Response
// @Failure		500		{object}	models.Response
// @Router			/messages/chat/{chat}/{limit} [get]
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

// @Summary		Update a message
// @Tags			messages
// @Description	Update a message with new content
// @ID				update-message
// @Accept			json
// @Produce		json
// @Param			message	body		models.Message	true	"Message object"
// @Success		200		{object}	models.Response
// @Failure		400		{object}	models.Response
// @Router			/message [put]
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

// @Summary		Delete a message
// @Tags			messages
// @Description	Delete a message with a specified ID
// @ID				delete-message
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Message ID"
// @Success		200	{object}	models.Response
// @Failure		400	{object}	models.Response
// @Router			/message/{id} [delete]
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

// FindPhoto godoc
//
//	@Summary		Find a photo
//	@Description	Find a photo based on the provided message data
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Param			messageData	body		models.Message	true	"Message data"
//	@Success		200			{object}	models.Response
//	@Failure		400			{object}	models.Response
//	@Failure		500			{object}	models.Response
//	@Router			/message/find-photo [post]
func (m *MessageHandler) FindPhoto(e echo.Context) error {
	var messageData models.Message
	if err := e.Bind(&messageData); err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params",
		})
	}

	id, err := m.usecase.FindPhoto(messageData)
	if err != nil && errors.Is(err, models.ErrBadParamInput) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Invalid params jfkls",
		})
	} else if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: mongo.ErrNoDocuments.Error(),
		})
	} else if err != nil {
		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: id,
	})
}

// @Summary		Uploads a photo
// @Tags			messages
// @Description	Uploads a photo with the specified timing, user ID, chatroom ID, and photo file
// @ID				uploadPhoto
// @Accept			multipart/form-data
// @Produce		json
// @Param			timing		formData	string	true	"Timing"
// @Param			user_id		formData	integer	true	"User ID"
// @Param			chatroom_id	formData	integer	true	"Chatroom ID"
// @Param			photo		formData	file	true	"Photo file"
// @Success		200			{object}	models.Response
// @Failure		400			{object}	models.Response
// @Failure		500			{object}	models.Response
// @Router			/upload-photo [post]
func (m *MessageHandler) UploadPhoto(e echo.Context) error {
	// ectracting data from form (time, user id, chatroom id and file)

	// converting time from string to time.Time
	timeStamp := e.FormValue("timing")
	sendTime, err := m.convertTimestamp(timeStamp)
	if err != nil {
		log.Println("here")
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: models.ErrBadParamInput.Error(),
		})
	}

	// converting user id to int
	user_id := e.FormValue("user_id")
	uid, err := strconv.Atoi(user_id)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Ivalid params",
		})
	}

	// converting chatroom id to int
	chatroom_id := e.FormValue("chatroom_id")
	cid, err := strconv.Atoi(chatroom_id)
	if err != nil {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: "Ivalid params",
		})
	}

	message := &models.Message{
		Sended:     sendTime,
		ChatroomID: cid,
		UserID:     uid,
	}

	uploaded, err := e.FormFile("photo")
	if err != nil {
		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Can not read uploaded file",
		})
	}

	insertedID, err := m.usecase.StorePhoto(*message)
	if err != nil {
		logger.STDLogger.Fatal(err.Error())
	}

	src, err := uploaded.Open()
	if err != nil {
		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Can not open uploaded file",
		})
	}
	defer src.Close()

	dst, err := os.Create("./static/images/" + insertedID)
	if err != nil {
		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Can not save uploaded file",
		})
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return e.JSON(500, models.Response{
			Message: "Failure",
			Content: "Can not write uploaded file",
		})
	}

	return e.JSON(200, models.Response{
		Message: "Success",
		Content: insertedID,
	})
}

// DeletePhoto godoc
//
//	@Summary		Delete a photo
//	@Description	Delete a photo based on the provided ID
//	@Tags			messages
//	@Param			id	path	string	true	"Photo ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Router			/message/delete-photo/{id} [delete]
func (m *MessageHandler) DeletePhoto(e echo.Context) error {
	id := e.Param("id")

	log.Println(id)
	deleted, err := m.usecase.DeletePhoto(id)
	if err != nil && errors.Is(err, models.ErrBadParamInput) {
		return e.JSON(400, models.Response{
			Message: "Failure",
			Content: err.Error(),
		})
	}

	if deleted != 0 {
		return e.JSON(200, models.Response{
			Message: "Success",
			Content: deleted,
		})
	} else {
		return e.JSON(200, models.Response{
			Message: "Failure",
			Content: deleted,
		})
	}
}

func (m *MessageHandler) convertTimestamp(timeString string) (time.Time, error) {
	log.Println(timeString, models.TimeLayout)
	t, err := time.Parse(models.TimeLayout, timeString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
