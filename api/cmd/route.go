package cmd

import (
	"api/cmd/auth"
	"api/cmd/feed"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	urlRoute := e.Group("/api/v1")
	urlRoute.POST("/auth", auth.Auth)
	urlRoute.POST("/feed", feed.PostFeed)
	urlRoute.GET("/feed", feed.GetFeed)
}
