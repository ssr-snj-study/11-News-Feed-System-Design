package repository

import (
	"gorm.io/gorm"
	"webserver/internal/device/entity"
)

type DeviceRepositoryImpl struct {
	DB *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepositoryImpl {
	return &DeviceRepositoryImpl{DB: db}
}

func (r *DeviceRepositoryImpl) Create(device *entity.Device) error {
	if err := r.DB.Create(&device).Error; err != nil {
		return err
	}
	return nil
}

func (r *DeviceRepositoryImpl) Update(device *entity.Device) error {
	if err := r.DB.Where("user_id = ?", device.UserId).Updates(device).Error; err != nil {
		return err
	}
	return nil
}

func (r *DeviceRepositoryImpl) GetByName(req *entity.Req) (*entity.Device, error) {
	user := &entity.User{}
	if res := r.DB.Where("name = ?", req.Name).Find(user); res.Error != nil {
		return 0, res.Error
	}
}
