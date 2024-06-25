package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restro-mgt/database"
	"restro-mgt/models"
)

func CreateTimeEntry(c *gin.Context) {
	var timeEntry models.TimeEntry
	if err := c.ShouldBindJSON(&timeEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&timeEntry)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, timeEntry)
}

func GetAllTimeEntries(c *gin.Context) {
	var timeEntries []models.TimeEntry
	result := database.DB.Find(&timeEntries)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, timeEntries)
}
