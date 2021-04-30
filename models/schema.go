package models

import (
	"log"
	"time"

	"github.com/DanielUlises98/mytelebot/KEYS"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Username  string
	Animes    []*Anime `gorm:"many2many:user_animes;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Anime struct {
	ID             string `gorm:"primaryKey"`
	Episodes       uint
	CurrentEpisode uint
	Name           string
	ImageMedium    string
	ImageOriginal  string
	StartDate      string
	EndDate        string
	Status         bool
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	User           []*User `gorm:"many2many:user_animes;"`
}

type UserAnimes struct {
	UserID     string `gorm:"foreignKey"`
	AnimeID    string `gorm:"foreignKey"`
	HourRemind string
	WeekDay    string
	RemindUser bool `gorm:"default:false"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
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
