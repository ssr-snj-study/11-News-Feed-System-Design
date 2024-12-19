package routes

import (
	"github.com/labstack/echo/v4"
	"webserver/cmd/app"
)

func RegisterRoutes(e *echo.Echo, deps *app.Dependencies) {
	// Auth routes
	e.POST("/auth", deps.AuthHandler.Auth)

	//// Feed routes
	//api := e.Group("/api/v1")
	//api.POST("/feed", deps.FeedHandler.PostFeed)
	//api.GET("/feed", deps.FeedHandler.GetFeed)
}
