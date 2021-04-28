package services

import (
	"github.com/gorilla/websocket"
	"go-scrum/models"
	"gorm.io/gorm"
	"log"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db,
	}
}

func (a *AuthService) Login(conn *websocket.Conn, auth *models.Auth) {
	err := conn.WriteJSON(auth)
	if err != nil {
		err := conn.Close()
		if err != nil {
			log.Println("Cannot close client connection")
		}
	}
}
func (a *AuthService) Logout(conn *websocket.Conn) {
	log.Printf("User with conn %s has logout\n", conn.RemoteAddr().String())
	err := conn.Close()
	if err != nil {
		return
	}
}
