package main

import (
	"log"
	"webserver/cmd/app"
	"webserver/cmd/routes"
)

func main() {
	// Initialize application
	application := app.InitializeApp()
	//defer application.DB.Close()

	// Set up dependencies
	dependencies := app.InitializeDependencies(application)

	// Register routes
	routes.RegisterRoutes(application.Echo, dependencies)

	// Start server
	log.Println("Server is running on port 1323")
	application.Echo.Logger.Fatal(application.Echo.Start(":1323"))
}
