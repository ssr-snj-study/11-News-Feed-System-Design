package entity

import "time"

type Device struct {
	Id             int       `json:"id"`
	DeviceToken    string    `json:"device_token"`
	UserId         int       `json:"user_id"`
	LastLoggedInAt time.Time `json:"last_logged_in_at"`
}

func (Device) TableName() string {
	return "device"
}
