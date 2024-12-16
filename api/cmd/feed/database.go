package feed

import (
	"api/config"
	"api/model"
	"errors"
	"time"
)

func createFeed(data *req) error {
	db := config.DB()
	posting := &model.Posting{
		UserId:      data.UserId,
		Contents:    data.Contents,
		Likes:       0,
		CreatedTime: time.Now(),
	}
	if err := db.Create(&posting).Error; err != nil {
		return err
	}
	return nil
}

func getFeed(data *req) (*model.Posting, error) {
	db := config.DB()
	posting := &model.Posting{}
	if res := db.Where("id = ?", data.FeedId).Find(posting); res.Error != nil {
		return nil, res.Error
	}
	if posting.UserId == 0 {
		return nil, errors.New("post is not exist")
	}
	return posting, nil

}

func getFollowerToken(userId int) error {
	db := config.DB()
	posting := &model.Posting{}
	query := `SELECT
	f.user_id AS follower_id,
		u.followed_user_id,
		d.device_token
	FROM
	follower f
	CROSS JOIN LATERAL
	unnest(f.followers) AS u(followed_user_id)
	LEFT JOIN
	device d
	ON
	u.followed_user_id = d.user_id
	where ? = Any(f.followers)`
	if res := db.Raw(query, userId).Find(posting); res.Error != nil {
		return res.Error
	}
	return nil
}
