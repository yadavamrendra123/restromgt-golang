package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restro-mgt/database"
	"restro-mgt/models"
	"restro-mgt/types"
)

func CreateCustomFormat(c *gin.Context) {
	var customFormat types.CustomFormat
	if err := c.ShouldBindJSON(&customFormat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customFormatModel := models.ConvertFromCustomFormat(customFormat)
	result := database.DB.Create(&customFormatModel)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, customFormatModel)
}

func GetAllCustomFormats(c *gin.Context) {
	var customFormats []models.CustomFormatModel
	result := database.DB.Find(&customFormats)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var response []types.CustomFormat
	for _, customFormatModel := range customFormats {
		customFormat, err := customFormatModel.ConvertToCustomFormat()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response = append(response, customFormat)
	}

	c.JSON(http.StatusOK, response)
}
