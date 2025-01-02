package repository

import "webserver/internal/feed/entity"

type FeedRepository interface {
	PostFeed(req *entity.Req) (int, error)
	GetFeed(req *entity.Req) (*entity.Posting, error)
}
