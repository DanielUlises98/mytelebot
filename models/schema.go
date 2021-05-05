package models

import (
	"fmt"
	"log"
	"strings"
	"time"

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

func InitDB(dsn string) *gorm.DB {
	fmt.Println("Setting up database")
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
	fmt.Println("Finished setting up the database")
	return db
}

func UrlToDsn(url string) string {
	character := []string{":", "@", ":", "/"}
	dsn := make([]string, 5)
	url = url[11:]
	for i, item := range character {
		p := strings.Index(url, item)
		dsn[i] = url[0:p]
		url = url[p+1:]
	}
	dsn[4] = url
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", dsn[0], dsn[1], dsn[2], dsn[3], dsn[4], "require")
}
