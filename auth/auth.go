package auth

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Bronku/iroon/crypto"
	"github.com/Bronku/iroon/models"
	"gorm.io/gorm"
)

type Config struct {
	SessionExpiration time.Duration
}

type Authenticator struct {
	db     *gorm.DB
	config Config
}

func New(db *gorm.DB, config Config) Authenticator {
	err := db.AutoMigrate(&models.User{}, &models.Token{})
	if err != nil {
		log.Fatal("failed to initialize authenticator", err)
	}
	return Authenticator{db, config}
}

var ErrUserNotFound = errors.New("user with this login doesn't exist")
var ErrWrongCredentials = errors.New("wrong credentials")

// #todo only used once, maybe remove it entirely later
func (a *Authenticator) verifyCredentials(login, password string) error {
	var user models.User

	result := a.db.First(&user, "login = ?", login)
	if result.Error != nil {
		return ErrUserNotFound
	}

	hash := crypto.PasswordHash(password, user.Salt)
	if hash == user.Password {
		return nil
	}

	return ErrWrongCredentials
}

func (a *Authenticator) Middleware(inner http.Handler) http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /login", getLogin)
	handler.HandleFunc("POST /login", a.login)
	handler.HandleFunc("GET /logout", a.logout)
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := a.getSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		inner.ServeHTTP(w, r)
	})

	return handler
}

var ErrUsernameTaken = errors.New("username taken")

func (a *Authenticator) AddUser(login, password string) error {
	var user models.User

	result := a.db.First(&user, "login = ?", login)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ErrUsernameTaken
	}

	salt := crypto.GenerateKey()
	hash := crypto.PasswordHash(password, salt)

	result = a.db.Save(&models.User{Login: login, Password: hash, Salt: salt})

	return result.Error
}
