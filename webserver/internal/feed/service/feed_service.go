package service

import (
	"webserver/internal/feed/entity"
	"webserver/internal/feed/repository"
)

type FeedService struct {
	Repo repository.FeedRepository
}

func (s *FeedService) PostingFeed(req *entity.Req) (int, error) {
	postNo, err := s.Repo.PostFeed(req)
	if err != nil {
		return 0, err
	}
	return postNo, nil

}

func (s *FeedService) GetFeed(req *entity.Req) (*entity.Posting, error) {
	return s.Repo.GetFeed(req)
}
