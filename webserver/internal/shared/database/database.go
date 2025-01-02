package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"webserver/internal/shared/config"
)

func NewDB(config config.DBConfig) (*gorm.DB, error) {
	connectInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Host, config.User, config.Password, config.DBName, config.Port)
	db, e := gorm.Open(postgres.Open(connectInfo), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	return db, e
}
