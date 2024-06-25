package models

import (
	"gorm.io/gorm"
	"restro-mgt/types"
)

type Event struct {
	gorm.Model
	Name      string         `json:"name"`
	EventDate types.DateOnly `json:"event_date"`
}
