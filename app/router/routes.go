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
			"message": "Hello dear!",
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/restaurants", controllers.CreateRestaurant)
		v1.GET("/restaurants", controllers.GetAllRestaurants)
		v1.GET("/restaurants/:id", controllers.GetRestaurantByID)
		v1.PUT("/restaurants/:id", controllers.UpdateRestaurant)
		v1.DELETE("/restaurants/:id", controllers.DeleteRestaurant)
	}

	return router
}
