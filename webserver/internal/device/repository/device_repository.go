package repository

import "webserver/internal/device/entity"

type DeviceRepository interface {
	Create(device *entity.Device) error
	Update(device *entity.Device) (int, error)
	GetByName(device *entity.Req) (*entity.Device, error)
}
