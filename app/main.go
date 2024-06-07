package main

import (
	"log"
	"restro-mgt/database"
	"restro-mgt/router"
)

func main() {
	database.InitDB()

	// Set up the router
	route := router.SetupRouter()

	// Run the server
	err := route.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
