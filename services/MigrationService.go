package services

import (
	"go-scrum/models"
	"gorm.io/gorm"
)

type MigrationService struct {
	db *gorm.DB
}

func NewMigrationService(db *gorm.DB) *MigrationService {
	return &MigrationService{
		db: db,
	}
}
func (m *MigrationService) Migrate() {
	err := m.db.AutoMigrate(&models.Room{})
	if err != nil {
		panic("Error trying to migrate")
	}
}
