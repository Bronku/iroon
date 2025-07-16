package main

import (
	_ "embed"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Bronku/iroon/auth"
	"github.com/Bronku/iroon/logging"
	"github.com/Bronku/iroon/models"
	"github.com/Bronku/iroon/server"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type config struct {
	Database struct {
		File string
	}

	Authentication struct {
		auth.Config
		DefaultLogin    string
		DefaultPassword string
	}

	Server struct {
		Addr              string
		ReadHeaderTimeout time.Duration
	}
}

//go:embed config.default.toml
var defaultConfig string

func main() {
	// load config
	var conf config
	_, err := toml.Decode(defaultConfig, &conf)
	if err != nil {
		log.Fatal(err)
	}
	_, err = toml.DecodeFile("config.toml", &conf)
	if err != nil {
		log.Println("using default config")
	}

	// load database
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true},
	)
	db, err := gorm.Open(sqlite.Open(conf.Database.File), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	err = db.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Product{})
	if err != nil {
		log.Fatal("couldn't migrate database")
	}

	// load auth
	authenticator := auth.New(db, conf.Authentication.Config)
	err = authenticator.AddUser(conf.Authentication.DefaultLogin, conf.Authentication.DefaultPassword)
	if err != nil && !errors.Is(err, auth.ErrUsernameTaken) {
		log.Println("add user:", err)
	}

	// load server
	h := server.New(db)
	var handler http.Handler = h
	handler = authenticator.Middleware(handler)
	handler = logging.Middleware(handler)
	log.Println("starting server")

	// start server
	server := &http.Server{
		Handler:           handler,
		Addr:              conf.Server.Addr,
		ReadHeaderTimeout: conf.Server.ReadHeaderTimeout,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
