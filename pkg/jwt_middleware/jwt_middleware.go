package jwtmiddleware

import (
	"chat/models"
	"chat/pkg/config"
	"fmt"
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
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
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
		if len(authHeader) == 0 {
			return e.JSON(401, models.Response{
				Message: "Failure",
				Content: models.ErrPermisionDenied.Error(),
			})
		}
		hdr := authHeader[0]
		tokenString := hdr[len("Bearer "):]

		if len(tokenString) == 0 {
			return e.JSON(401, models.Response{
				Message: "Failure",
				Content: models.ErrPermisionDenied.Error(),
			})
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		})
		if err != nil {
			fmt.Println(err.Error())
			return e.JSON(401, models.Response{
				Message: "Failure",
				Content: models.ErrPermisionDenied.Error(),
			})
		}

		if !token.Valid {
			return e.JSON(401, models.Response{
				Message: "Failure",
				Content: models.ErrPermisionDenied.Error(),
			})
		} else {
			next(e)
		}

		return nil
	}
}
