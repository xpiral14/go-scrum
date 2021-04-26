package controllers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"go-scrum/constants"
	"go-scrum/models"
	"go-scrum/services"
	"log"
	"net/http"
	"strconv"
)

type WebSocketController struct{}

func NewWebSocketController() *WebSocketController {
	return &WebSocketController{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var authAttributes services.Auth
var authService = services.NewAuthService()

func (ws *WebSocketController) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}
		command, _ := strconv.Atoi(string(message[0:5]))
		payload := message[5:]

		switch command {
		case constants.HELLO_WORLD:
			{
				err := conn.WriteMessage(messageType, []byte("Seja bem vindo"))
				if err != nil {
					err := conn.Close()
					if err != nil {
						log.Fatalln("Some error occurend when trying to close connection: " + err.Error())
					}
				}
			}
		case constants.LOGIN:
			{
				err := json.Unmarshal(payload, &authAttributes)
				if err != nil {
					err := conn.WriteJSON(models.Error{
						Code:    1,
						Name:    "Invalid JSON",
						Message: "JSON has invalid patter or is not accepted",
					})
					if err != nil {
						log.Println(err)
					}
				}
				authService.Login(conn, &authAttributes)
			}
		case constants.LOGOUT:
			{
				authService.Logout(conn)
			}
		default:
			conn.WriteMessage(messageType, []byte("Comando inválido"))
		}
	}
}
