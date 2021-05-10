package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DanielUlises98/mytelebot/models"
	"github.com/joho/godotenv"
)

var (
	port, publicUrl, token, dsn string
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port = os.Getenv("PORT")
	publicUrl = os.Getenv("PUBLIC_URL")
	token = os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("Error when loading the telegram token")
	}
	dsn = os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("LOCAL_DSN")
	} else {
		dsn = models.UrlToDsn(dsn)
	}
}

func main() {
	// InitEnv()
	// db := models.InitDB(dsn)
	// bot := tbBot.StartBot(token, port, publicUrl)
	// tbBot.InitHandlers(db, bot)
	// reminder.Init(db, bot)
	// bot.Start()
	// PARSE TIMEZONE TO UTC TIME AN WORK ON THAT
	load, _ := time.LoadLocation("America/Mexico_City")
	nowCurrent := time.Now().In(load)
	weekday := nowCurrent.Weekday()
	fmt.Println(weekday)
	t := time.Date(nowCurrent.Year(), nowCurrent.Month(), nowCurrent.Day()+6-1, 17, 30, 0, 0, load)
	fmt.Println(t)
	fmt.Println(nowCurrent.Sub(t))
}

/*
1.-The bot will remind you to watch your anime on the day is published
2.-You cand add the anime that you to be reminded to you

1.-The list of animes will contain if is being published or has been finished
2.-if the anime has alredy been finished, you can choose a day to be reminded
*/
