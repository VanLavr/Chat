package jwtmiddleware

import (
	"chat/models"
	"chat/pkg/config"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type JwtMiddleware struct {
	secret string
}

func NewJwtMiddlware() *JwtMiddleware {
	secret := config.Con.GetSecret()
	return &JwtMiddleware{secret: secret}
}

type jwtClaims struct {
	User models.User
	jwt.StandardClaims
}

func (j *JwtMiddleware) GenerateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			Issuer:    user.Name,
		},
	})

	stringifiedToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		log.Fatal(err)
	}

	return stringifiedToken
}

func (j *JwtMiddleware) ValidateToken(next func(e echo.Context) error) func(e echo.Context) error {
	return func(e echo.Context) error {
		header := e.Request().Header
		authHeader := header["Authorization"]
		tokenString := authHeader[len("Bearer "):]

		if len(tokenString) == 0 {
			return e.JSON(401, models.Response{
				Message: "Failure",
				Content: models.ErrPermisionDenied,
			})
		}

		token, err := jwt.ParseWithClaims(tokenString[0], &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		})
		if err != nil {
			return e.JSON(401, models.Response{
				Message: "Failure",
				Content: models.ErrPermisionDenied,
			})
		}

		if !token.Valid {
			return e.JSON(401, models.Response{
				Message: "Failure",
				Content: models.ErrPermisionDenied,
			})
		} else {
			next(e)
		}

		return nil
	}
}
