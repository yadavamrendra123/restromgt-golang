package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restro-mgt/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connected to the database!",
		})
	})

	router.GET("/restaurants", controller.GetRestaurants)

	return router
}
