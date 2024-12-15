package cmd

import (
	"api/cmd/auth"
	"api/cmd/device"
	"api/cmd/feed"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	e.POST("/auth", auth.Auth)
	e.POST("/SetDevice", device.UpsertDevice)

	urlRoute := e.Group("/api/v1")
	urlRoute.Use(echojwt.WithConfig(echojwt.Config{SigningKey: []byte("test"), TokenLookup: "cookie:access-token"}))
	urlRoute.POST("/feed", feed.PostFeed)
	urlRoute.GET("/feed", feed.GetFeed)

}
