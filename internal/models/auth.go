package models

import "time"

type Token struct {
	User       string
	Expiration time.Time
}
