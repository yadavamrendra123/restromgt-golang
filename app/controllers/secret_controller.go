package controllers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"restro-mgt/models"
)

type SecretController struct {
	DB *gorm.DB
}

func (sc *SecretController) CreateSecret(c *gin.Context) {
	var secret models.Secret
	if err := c.ShouldBindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Encrypt the text before saving to the database
	encryptedText, err := encryptString(secret.EncryptedText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt secret"})
		return
	}
	secret.EncryptedText = encryptedText

	// Create the secret in the database
	if err := sc.DB.Create(&secret).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, secret)
}

// encryptString encrypts a string using base64 encoding.
func encryptString(input string) (string, error) {
	encrypted := base64.StdEncoding.EncodeToString([]byte(input))
	return encrypted, nil
}

func (sc *SecretController) GetAllSecrets(c *gin.Context) {
	var secrets []models.Secret
	if err := sc.DB.Find(&secrets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var secretResponses []models.SecretResponse
	for _, secret := range secrets {
		decryptedText := decryptString(secret.EncryptedText)
		secretResponses = append(secretResponses, models.SecretResponse{
			ID:            secret.ID,
			CreatedAt:     secret.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:     secret.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Name:          secret.Name,
			DecryptedText: decryptedText,
		})
	}

	c.JSON(http.StatusOK, secretResponses)
}

// decryptString decrypts a base64-encoded string.
func decryptString(input string) string {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Printf("Error decoding base64: %v", err)
		return ""
	}
	return string(decoded)
}
