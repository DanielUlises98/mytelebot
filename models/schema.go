package models

import (
	"log"
	"time"

	"github.com/DanielUlises98/mytelebot/KEYS"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"autoIncrement;primaryKey"`
	Username  string
	ChatId    string
	Animes    []*Anime `gorm:"many2many:user_animes;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Anime struct {
	ID      uint `gorm:"autoIncrement;primaryKey"`
	IdAnime string
	User    []*User `gorm:"many2many:user_animes;"`
}

type UserAnimes struct {
	ID        uint `gorm:"autoIncrement"`
	UserID    uint `gorm:"foreignKey"`
	AnimeID   uint `gorm:"foreignKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func InitDB() *gorm.DB {

	dsn := KEYS.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error(), "Un error al conectar con la base de datos")
	}

	/*if !db.Migrator().HasTable(&User{}) {
		db.Migrator().CreateTable(&User{})
	}*/
	err = db.SetupJoinTable(&User{}, "Animes", &UserAnimes{})
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.SetupJoinTable(&Anime{}, "User", &UserAnimes{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}
