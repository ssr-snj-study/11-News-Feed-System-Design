package routes

import (
	"github.com/labstack/echo/v4"
	"webserver/cmd/app"
	"webserver/internal/shared/middleware"
)

func RegisterRoutes(e *echo.Echo, deps *app.Dependencies) {
	// Auth routes
	e.POST("/auth", deps.AuthHandler.Auth)
	e.POST("/SetDevice", deps.DeviceHandler.Device)

	//// Feed routes
	api := e.Group("/api/v1")
	api.Use(middleware.JWTMiddleware())
	api.POST("/feed", deps.FeedHandler.PostFeed)
	api.GET("/feed", deps.FeedHandler.GetFeed)
}
