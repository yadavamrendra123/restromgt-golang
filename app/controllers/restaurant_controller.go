package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurants(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"restaurants": "Restaurants list ooo lal la",
	})
}
