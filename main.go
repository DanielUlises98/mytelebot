package main

import (
	"log"
	"os"

	"github.com/DanielUlises98/mytelebot/models"
	"github.com/DanielUlises98/mytelebot/reminder"
	"github.com/DanielUlises98/mytelebot/tbBot"
	"github.com/joho/godotenv"
)

var (
	port, publicUrl, token, dsn string
)

func Init() {
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
		dsn = os.Getenv("DSN")
	} else {
		dsn = models.UrlToDsn(dsn)
	}
}

func main() {

	db := models.InitDB(dsn)
	bot := tbBot.StartBot(token, port, publicUrl)
	tbBot.InitHandlers(db, bot)
	reminder.Init(db, bot)
	bot.Start()
}

/*
1.-The bot will remind you to watch your anime on the day is published
2.-You cand add the anime that you to be reminded to you

1.-The list of animes will contain if is being published or has been finished
2.-if the anime has alredy been finished, you can choose a day to be reminded
*/
