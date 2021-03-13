package models

import (
	"log"

	"github.com/DanielUlises98/mytelebot/KEYS"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := KEYS.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error(), " Un error al conectar con la base de datos")
	}
	return db
}
