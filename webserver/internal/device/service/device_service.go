package service

import (
	"errors"
	"webserver/internal/device/entity"
	"webserver/internal/device/repository"
)

type DeviceService struct {
	Repo repository.DeviceRepository
}

func (s *DeviceService) UpsertDevice(req *entity.Req) (int, error) {
	var err error
	device := s.Repo.GetByName(req)
	userId := 0

	switch {
	case device.Id == 0:
		return 0, errors.New("you have to register")
	case device.DeviceToken == "":
		userId, err = s.Repo.Create(req)
		if err != nil {
			return 0, err
		}
	case device.DeviceToken != req.DeviceToken:
		userId, err = s.Repo.Update(req)
		if err != nil {
			return 0, err
		}
	}

	return userId, nil
}
