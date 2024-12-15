package auth

import (
	"api/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Auth(c echo.Context) error {
	b := new(model.User)
	if err := c.Bind(b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	userId, err := CheckUserName(b.Name)
	if err != nil {
		c.Logger().Warn(err)
	}

	accessToken, err := createJWT(userId)
	if err != nil {
		return echo.ErrInternalServerError
	}

	cookie := new(http.Cookie)
	cookie.Name = "access-token"
	cookie.Value = accessToken
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24)

	c.SetCookie(cookie)
	return nil
}

func createJWT(userId int) (string, error) {
	mySigningKey := []byte("test")

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tk, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tk, nil
}
