package service

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
	"webserver/internal/auth/entity"
	"webserver/internal/auth/repository"
)

type AuthService struct {
	Repo repository.AuthRepository
}

func (s *AuthService) Authenticate(req *entity.Req) (*http.Cookie, error) {
	user := new(entity.User)
	userId, err := s.Repo.AuthCheck(req.Name, user)
	if err != nil {
		return nil, err
	}
	accessToken, err := createJWT(userId)

	if err != nil {
		return nil, err
	}

	cookie := new(http.Cookie)
	cookie.Name = "access-token"
	cookie.Value = accessToken
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24)

	return cookie, nil
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
