package app

import (
	"webserver/internal/auth/handler"
	"webserver/internal/auth/repository"
	"webserver/internal/auth/service"
)

type Dependencies struct {
	AuthHandler handler.AuthHandler
}

func InitializeDependencies(app *App) *Dependencies {

	// Auth dependencies
	authRepo := repository.NewAuthRepository(app.DB)
	authService := service.AuthService{authRepo}
	authHandler := handler.AuthHandler{AuthService: &authService}

	return &Dependencies{
		AuthHandler: authHandler,
	}
}
