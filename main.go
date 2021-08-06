package main

import (
	"fmt"
	"go-scrum/controllers"
	"go-scrum/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	port := 8081
	db, err := gorm.Open(sqlite.Open("tmp.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	services.StartServices(db)
	services.MigrationServiceInstace.Migrate()

	http.HandleFunc("/ws", controllers.NewWebSocketController().Handler)

	log.Printf("Server running on localhost:%v", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatalln(err)
	}
}
