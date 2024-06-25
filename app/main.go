package main

import (
	"log"
	"restro-mgt/controllers"
	"restro-mgt/database"
	"restro-mgt/router"
)

func main() {
	db := database.InitDB()

	// Create an instance of SecretController
	secretController := controllers.SecretController{DB: db}

	// Set up the router with SecretController
	route := router.SetupRouter(secretController)

	// Run the server
	err := route.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
