package auth

import (
	_ "embed"
	"errors"
	"log"
	"net/http"

	"github.com/Bronku/iroon/crypto"
	"github.com/Bronku/iroon/models"
	"gorm.io/gorm"
)

type Authenticator struct {
	db *gorm.DB
}

func New(db *gorm.DB) Authenticator {
	err := db.AutoMigrate(&models.User{}, &models.Token{})
	if err != nil {
		log.Fatal("failed to initialize authenticator", err)
	}
	return Authenticator{db}
}

// #todo only used once, maybe remove it entierly later
func (a *Authenticator) verifyCredentials(login, password string) error {
	var user models.User

	result := a.db.First(&user, "login = ?", login)
	if result.Error != nil {
		return errors.New("user with this login doesn't exist")
	}

	hash := crypto.PasswordHash(password, user.Salt)
	if hash == user.Password {
		return nil
	}
	return errors.New("wrong credentials")
}

func (a *Authenticator) Middleware(in http.Handler) http.Handler {
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
		in.ServeHTTP(w, r)
	})
	return handler
}

//	func (s *Store) AddUser(login, password string) error {
//		_, exists := s.GetUser(login)
//		if exists {
//			return errors.New("the user already exists")
//		}
//		query := "insert into user (login, password, salt) values(?, ?, ?)"
//		salt := crypto.GenerateKey()
//		hash := crypto.PasswordHash(password, salt)
//		_, err := s.db.Exec(query, login, hash, salt)
//		if err == nil {
//			s.users[login] = models.User{Password: hash, Salt: salt}
//		}
//		return err
//	}
func (a *Authenticator) AddUser(login, password string) error {
	var user models.User

	result := a.db.First(&user, "login = ?", login)
	if result.Error != gorm.ErrRecordNotFound {
		return errors.New("user already exists")
	}

	salt := crypto.GenerateKey()
	hash := crypto.PasswordHash(password, salt)

	result = a.db.Save(&models.User{Login: login, Password: hash, Salt: salt})

	return result.Error
}
