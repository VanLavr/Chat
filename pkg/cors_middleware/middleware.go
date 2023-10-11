package corsmiddleware

import (
	"chat/pkg/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Middleware struct {
	secret string
}

func NewMiddleware(e *echo.Echo) {
	originsList := config.Con.GetOrigin()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{originsList},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		MaxAge:           3600,
	}))
}
