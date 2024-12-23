package repository

import "webserver/internal/device/entity"

type DeviceRepository interface {
	Create(device *entity.Req) (int, error)
	Update(device *entity.Req) (int, error)
	GetByName(device *entity.Req) *entity.Device
}
