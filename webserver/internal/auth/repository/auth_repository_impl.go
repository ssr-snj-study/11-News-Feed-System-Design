package repository

import (
	"gorm.io/gorm"
	"webserver/internal/auth/entity"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{DB: db}
}

func (r *AuthRepositoryImpl) AuthCheck(userName string, user *entity.User) (int, error) {
	if res := r.DB.Where("name = ?", userName).Find(user); res.Error != nil {
		return 0, res.Error
	}
	return user.Id, nil
}
