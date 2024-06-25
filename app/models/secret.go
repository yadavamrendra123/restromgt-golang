package models

import (
	"gorm.io/gorm"
	"time"
)

type Secret struct {
	ID            uint           `gorm:"primaryKey" json:"ID"`
	CreatedAt     time.Time      `json:"CreatedAt"`
	UpdatedAt     time.Time      `json:"UpdatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Name          string         `json:"name"`
	EncryptedText string         `gorm:"column:encrypted_text;type:text" json:"text_to_be_encrypted"`
}

type SecretResponse struct {
	ID            uint   `json:"ID"`
	CreatedAt     string `json:"CreatedAt"`
	UpdatedAt     string `json:"UpdatedAt"`
	Name          string `json:"name"`
	DecryptedText string `json:"decrypted_text"`
}
