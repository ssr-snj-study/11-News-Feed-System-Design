package entity

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

type Req struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
