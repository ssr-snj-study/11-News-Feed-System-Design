package app

import (
	authHandler "webserver/internal/auth/handler"
	authRep "webserver/internal/auth/repository"
	authService "webserver/internal/auth/service"
	deviceHandler "webserver/internal/device/handler"
	deviceRep "webserver/internal/device/repository"
	deviceService "webserver/internal/device/service"
)

type Dependencies struct {
	AuthHandler   authHandler.AuthHandler
	DeviceHandler deviceHandler.DeviceHandler
}

func InitializeDependencies(app *App) *Dependencies {

	// Auth dependencies
	authRepo := authRep.NewAuthRepository(app.DB)
	authService := authService.AuthService{Repo: authRepo}
	authHandler := authHandler.AuthHandler{AuthService: &authService}

	deviceRepo := deviceRep.NewDeviceRepository(app.DB)
	deviceService := deviceService.DeviceService{Repo: deviceRepo}
	deviceHandler := deviceHandler.DeviceHandler{DeviceService: &deviceService}

	return &Dependencies{
		AuthHandler:   authHandler,
		DeviceHandler: deviceHandler,
	}
}
