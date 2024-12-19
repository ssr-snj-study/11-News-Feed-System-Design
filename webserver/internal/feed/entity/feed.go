package entity

import "time"

type Posting struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Contents    string    `json:"contents"`
	Likes       int       `json:"likes"`
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

type Follower struct {
	Id        int   `json:"id"`
	UserId    int   `json:"user_id"`
	Followers []int `json:"followers"`
}

func (Follower) TableName() string {
	return "follower"
}
