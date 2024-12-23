package repository

import (
	"gorm.io/gorm"
	"time"
	"webserver/internal/device/entity"
)

type DeviceRepositoryImpl struct {
	DB *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepositoryImpl {
	return &DeviceRepositoryImpl{DB: db}
}

func (r *DeviceRepositoryImpl) Create(req *entity.Req) (int, error) {
	device := &entity.Device{
		LastLoggedInAt: time.Now(),
		UserId:         req.Id,
		DeviceToken:    req.DeviceToken,
	}
	if err := r.DB.Create(&device).Scan(device).Error; err != nil {
		return 0, err
	}
	return device.Id, nil
}

func (r *DeviceRepositoryImpl) Update(req *entity.Req) (int, error) {
	device := &entity.Device{
		LastLoggedInAt: time.Now(),
		UserId:         req.Id,
		DeviceToken:    req.DeviceToken,
	}
	if err := r.DB.Where("user_id = ?", device.UserId).Updates(device).Scan(device).Error; err != nil {
		return 0, err
	}
	return device.UserId, nil
}

func (r *DeviceRepositoryImpl) GetByName(req *entity.Req) *entity.Device {
	device := &entity.Device{}
	r.DB.Table("device").Select("device.device_token, user_tb.id as user_id").Joins("join user_tb on device.user_id = user_tb.id").Where("user_tb.name = ?", req.Name).Scan(device)
	return device
}
