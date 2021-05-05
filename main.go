package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DanielUlises98/mytelebot/models"
	"github.com/DanielUlises98/mytelebot/reminder"
	"github.com/DanielUlises98/mytelebot/tbBot"
	"github.com/joho/godotenv"
)

// var (
// 	api API.DBClient
// )

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Error when loading the telegram token")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("DSN")
	} else {
		dsn = models.UrlToDsn(dsn)
	}

	db := models.InitDB(dsn)
	bot := tbBot.StartBot(token)
	tbBot.InitHandlers(db, bot)
	reminder.Init(db, bot)
	fmt.Println("Starting bot")
	bot.Start()
	//urlToDsn("postgres://cfercjojdpxdbn:a1bc0cc912b0c9652a0a4c3969dbd6873c0fb8ecc6c456f95da4963daee88fdd@ec2-52-87-107-83.compute-1.amazonaws.com:5432/dl0llkn4jk1ki")
}

/*
Make the api out of kitsyu

//BASE API OF KITSU https://kitsu.io/api/edge


1.-The bot will remind you to watch your anime on the day is published
2.-You cand add the anime that you to be reminded to you

1.-The list of animes will contain if is being published or has been finished
2.-if the anime has alredy been finished, you can choose a day to be reminded
*/
