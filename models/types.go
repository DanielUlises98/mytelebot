package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"autoIncrement"`
	Username  string
	ChatId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DBTLClient struct {
	DB *gorm.DB
}
