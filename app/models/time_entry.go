package models

import (
	"gorm.io/gorm"
	"restro-mgt/types"
)

type TimeEntry struct {
	gorm.Model
	Name string         `json:"name"`
	Time types.TimeOnly `json:"time"`
}
