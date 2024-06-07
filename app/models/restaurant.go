package models

import (
	"time"

	"gorm.io/gorm"
)

type Restaurant struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Address     string
	PhoneNumber string
	Website     string
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
