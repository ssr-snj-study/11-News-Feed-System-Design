package repository

import (
	"webserver/internal/auth/entity"
)

type AuthRepository interface {
	AuthCheck(userName string, user *entity.User) (int, error)
}
