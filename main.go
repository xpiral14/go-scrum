package main

import (
	"go-scrum/controllers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var Db *gorm.DB

func main() {
	db, err := gorm.Open(sqlite.Open("tmp.db"), &gorm.Config{})
	Db = db
	if err != nil {
		panic("Failed to connect database")
	}

	http.HandleFunc("/ws", controllers.NewWebSocketController().Handler)

	log.Println("Server running on localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
