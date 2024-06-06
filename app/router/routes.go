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
			"message": "Hello babe!",
		})
	})

	router.GET("/restaurants", controller.GetRestaurants)

	return router
}
