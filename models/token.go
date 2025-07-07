package models

import "time"

type Token struct {
	Token      string `gorm:"primarykey"`
	User       string
	Expiration time.Time
}
