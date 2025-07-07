package main

import (
	"log"
	"net/http"

	"github.com/Bronku/iroon/logging"
	"github.com/Bronku/iroon/models"
	"github.com/Bronku/iroon/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Product{})

	//s.AddUser("admin", "secret")
	h := server.New(db)

	var handler http.Handler = h
	handler = logging.Middleware(handler)
	//handler = auth.New(s).Middleware(handler)
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
