package model

import "time"

type User struct {
	Id          int       `json:"id"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"created_time"`
}

func (User) TableName() string {
	return "user_tb"
}
