package model

import "time"

type User struct {
	Id          int       `json:"id"`
	UserId      string    `json:"user_id"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"created_time"`
}

func (User) TableName() string {
	return "user_tb"
}

type Device struct {
	Id             int       `json:"id"`
	DeviceToken    string    `json:"device_token"`
	UserId         int       `json:"user_id"`
	LastLoggedInAt time.Time `json:"last_logged_in_at"`
}

func (Device) TableName() string {
	return "device"
}

type Posting struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Contents    string    `json:"contents"`
	Like        int       `json:"like"`
	CreatedTime time.Time `json:"created_time"`
}

func (Posting) TableName() string {
	return "posting"
}

type Reply struct {
	Id          int       `json:"id"`
	PostingId   int       `json:"posting_id"`
	Contents    string    `json:"contents"`
	CreatedTime time.Time `json:"created_time"`
}

func (Reply) TableName() string {
	return "reply"
}