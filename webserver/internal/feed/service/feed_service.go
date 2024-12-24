package service

import "webserver/internal/feed/repository"

type FeedService struct {
	Repo *repository.FeedRepository
}
