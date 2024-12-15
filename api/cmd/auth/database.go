package auth

import (
	"api/config"
	"api/model"
)

func CheckUserName(userName string) (int, error) {
	db := config.DB()
	user := &model.User{}
	if res := db.Where("name = ?", userName).Find(user); res.Error != nil {
		return 0, res.Error
	}
	return user.Id, nil
}
