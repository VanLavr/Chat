package delivery

import (
	"chat/models"
	jwtmiddleware "chat/pkg/jwt_middleware"
)

type MessageHandler struct {
	usecase models.MessageUsecase
	jwtmiddleware.JwtMiddleware
}
