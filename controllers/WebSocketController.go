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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
		payload := message[3:]

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
				auth := models.Auth{}
				err := json.Unmarshal(payload, &auth)
				if err != nil {
					err := conn.WriteJSON(models.Error{
						Code:    1,
						Name:    "Invalid JSON",
						Message: "JSON has invalid patter or is not accepted",
					})
					if err != nil {
						message := services.ErrorServiceInstance.ToBytes(models.Error{
							Code:    constants.INVALID_PAYLOAD,
							Message: "Invalid payload",
							Name:    "INVALID_PAYLOAD",
						})
						_ = conn.WriteJSON(message)
					}
				}
				services.AuthServiceInstance.Login(conn, &auth)
			}
		case constants.LOGOUT:
			{
				services.AuthServiceInstance.Logout(conn)
			}
		case constants.CREATE_ROOM:
			{
				room := models.Room{}
				err := json.Unmarshal(payload, &room)
				if err != nil {
					message := services.ErrorServiceInstance.ToBytes(models.Error{
						Code:    constants.INVALID_PAYLOAD,
						Message: "Invalid payload",
						Name:    "INVALID_PAYLOAD",
					})
					_ = conn.WriteJSON(message)
				}
				services.RoomServiceInstance.CreateRoom(&room)
				_ = conn.WriteJSON(models.NewResponse(constants.CREATED_ROOM, room))
			}
		default:
			log.Println(command)
			if err := conn.WriteMessage(messageType, []byte("Comando inválido")); err != nil {
				_ = conn.Close()
			}
		}
	}
}
