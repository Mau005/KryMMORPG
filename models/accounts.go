package models

import (
	"time"

	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	NameAccount string
	Password    string
	secret      string
	Email       string
	TypeAcce    uint8
	PremEndDays time.Time
}
