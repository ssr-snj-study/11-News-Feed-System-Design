package device

import (
	"api/config"
	"api/model"
	"errors"
	"time"
)

func createDevice(data *params) error {
	db := config.DB()
	device := &model.Device{
		UserId:         data.UserId,
		DeviceToken:    data.Token,
		LastLoggedInAt: time.Now(),
	}
	if err := db.Create(&device).Error; err != nil {
		return err
	}
	return nil
}

func updateDevice(data *params) error {
	db := config.DB()
	device := &model.Device{
		DeviceToken:    data.Token,
		LastLoggedInAt: time.Now(),
	}
	if err := db.Where("user_id = ?", data.UserId).Updates(device).Error; err != nil {
		return err
	}
	return nil
}

func CheckDevice(data *params) (int, error) {
	db := config.DB()
	user := &model.User{}
	if res := db.Where("name = ?", data.Name).Find(user); res.Error != nil {
		return 0, res.Error
	}
	data.UserId = user.Id
	device := &model.Device{}
	if res := db.Where("user_id = ?", user.Id).Find(device); res.Error != nil {
		return 0, res.Error
	}
	switch device.DeviceToken {
	case "":
		return 1, nil
	case data.Token:
		return 0, errors.New("token exist")
	default:
		return 2, nil
	}
}
