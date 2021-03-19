package models

import (
	"log"

	"github.com/DanielUlises98/mytelebot/KEYS"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Animes struct {

// }

func InitDB() *gorm.DB {
	dsn := KEYS.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error(), " Un error al conectar con la base de datos")
	}

	/*if !db.Migrator().HasTable(&User{}) {
		db.Migrator().CreateTable(&User{})
	}*/
	db.AutoMigrate(&User{})
	return db
}
