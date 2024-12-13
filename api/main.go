package main

import (
	"api/cmd"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	cmd.Route(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
