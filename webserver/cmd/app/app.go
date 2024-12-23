package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"log"
	"webserver/internal/shared/config"
	"webserver/internal/shared/database"
)

type App struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func InitializeApp() *App {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := database.NewDB(cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Set up Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &App{
		Echo: e,
		DB:   db,
	}
}

func (app *App) Close() {
	// DB 연결 닫기
	sqlDB, err := app.DB.DB()
	if err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("error while closing database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
	// Echo 서버 종료(필요한 경우 추가 가능)
}
