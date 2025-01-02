package repository

import (
	"gorm.io/gorm"
	"time"
	"webserver/internal/feed/entity"
)

type FeedRepositoryImpl struct {
	DB *gorm.DB
}

func NewFeedRepository(db *gorm.DB) *FeedRepositoryImpl {
	return &FeedRepositoryImpl{DB: db}
}

func (r *FeedRepositoryImpl) PostFeed(req *entity.Req) (int, error) {
	feed := &entity.Posting{
		UserId:      req.UserId,
		Contents:    req.Contents,
		Likes:       0,
		CreatedTime: time.Now(),
	}
	if err := r.DB.Create(&feed).Scan(feed).Error; err != nil {
		return 0, err
	}
	return feed.Id, nil
}

func (r *FeedRepositoryImpl) GetFeed(req *entity.Req) (*entity.Posting, error) {
	posting := &entity.Posting{}
	if res := r.DB.Where("user_id = ?", req.UserId).Find(posting); res.Error != nil {
		return posting, res.Error
	}
	return posting, nil
}
