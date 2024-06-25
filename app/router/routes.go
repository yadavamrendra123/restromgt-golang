package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restro-mgt/controllers"
)

func SetupRouter(secretController controllers.SecretController) *gin.Engine {
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
		v1.POST("/events", controllers.CreateEvent)
		v1.GET("/events", controllers.GetAllEvents)
		v1.POST("/time-entries", controllers.CreateTimeEntry)
		v1.GET("/time-entries", controllers.GetAllTimeEntries)
		v1.POST("/secrets", secretController.CreateSecret) // Updated
		v1.GET("/secrets", secretController.GetAllSecrets) // Updated
		v1.POST("/custom-formats", controllers.CreateCustomFormat)
		v1.GET("/custom-formats", controllers.GetAllCustomFormats)
	}

	return router
}
