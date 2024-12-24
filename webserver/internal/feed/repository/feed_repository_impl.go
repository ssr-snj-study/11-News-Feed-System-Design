package repository

import (
	"gorm.io/gorm"
	"webserver/internal/feed/entity"
)

type FeedRepositoryImpl struct {
	DB *gorm.DB
}

func NewFeedRepository(db *gorm.DB) *FeedRepositoryImpl {
	return &FeedRepositoryImpl{DB: db}
}

func (r *FeedRepositoryImpl) PostFeed(req *entity.Req) error {
	return nil
}

func (r *FeedRepositoryImpl) GetFeed(req *entity.Req) (entity.Posting, error) {
	posting := entity.Posting{}

	return posting, nil
}
