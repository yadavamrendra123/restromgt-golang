package models

import (
	"gorm.io/gorm"
	"restro-mgt/types"
)

type CustomFormatModel struct {
	gorm.Model
	Value string `json:"value" gorm:"type:varchar(255)"`
}

// ConvertToCustomFormat converts the CustomFormatModel to CustomFormat.
func (m CustomFormatModel) ConvertToCustomFormat() (types.CustomFormat, error) {
	var cf types.CustomFormat
	err := cf.FromString(m.Value)
	return cf, err
}

// ConvertFromCustomFormat converts CustomFormat to CustomFormatModel.
func ConvertFromCustomFormat(cf types.CustomFormat) CustomFormatModel {
	return CustomFormatModel{Value: cf.ToString()}
}
