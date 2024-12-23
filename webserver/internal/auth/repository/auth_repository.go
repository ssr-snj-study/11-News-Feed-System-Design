package repository

import (
	"webserver/internal/shared/model"
)

type AuthRepository interface {
	AuthCheck(userName string, user *model.User) (int, error)
}
