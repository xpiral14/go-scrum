package services

import "gorm.io/gorm"

var AuthServiceInstance *AuthService
var RoomServiceInstance *RoomService
var ErrorServiceInstance *ErrorService
var MigrationServiceInstace *MigrationService

func StartServices(db *gorm.DB) {
	AuthServiceInstance = NewAuthService(db)
	RoomServiceInstance = NewRoomService(db)
	MigrationServiceInstace = NewMigrationService(db)
	ErrorServiceInstance = NewErrorService()
}
