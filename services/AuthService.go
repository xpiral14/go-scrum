package services

import (
	"github.com/gorilla/websocket"
	"log"
)

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Login(conn *websocket.Conn, auth *Auth) {
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
