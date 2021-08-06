package services

import (
	"go-scrum/models"
	"gorm.io/gorm"
)

type RoomService struct {
	db *gorm.DB
}

func NewRoomService(db *gorm.DB) *RoomService {
	return &RoomService{
		db,
	}
}

func (r *RoomService) CreateRoom(room *models.Room) *models.Room {
	r.db.Create(room)

	return room
}
