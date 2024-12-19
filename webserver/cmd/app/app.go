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
