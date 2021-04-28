package models

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}
