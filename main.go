package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Bronku/iroon/auth"
	"github.com/Bronku/iroon/logging"
	"github.com/Bronku/iroon/models"
	"github.com/Bronku/iroon/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true},
	)
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Product{})

	a := auth.New(db)
	err = a.AddUser("admin", "secret")
	if err != nil {
		log.Println("add user:", err)
	}

	h := server.New(db)

	var handler http.Handler = h
	handler = a.Middleware(handler)
	handler = logging.Middleware(handler)
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
