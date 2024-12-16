package feed

import (
	"api/config"
	"api/model"
)

func createFeed(data *req) error {
	db := config.DB()
	posting := &model.Posting{
		UserId:   data.UserId,
		Contents: data.Contents,
	}
	if err := db.Create(&posting).Error; err != nil {
		return err
	}
	return nil
}
