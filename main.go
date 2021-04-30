package main

import (
	"fmt"

	"github.com/DanielUlises98/mytelebot/models"
	"github.com/DanielUlises98/mytelebot/reminder"
	"github.com/DanielUlises98/mytelebot/tbBot"
)

// var (
// 	api API.DBClient
// )

func main() {
	db := models.InitDB()
	bot := tbBot.StartBot()
	tbBot.InitHandlers(db, bot)
	reminder.Init(db, bot)
	fmt.Println("Starting bot")
	bot.Start()
}

/*
Make the api out of kitsyu

//BASE API OF KITSU https://kitsu.io/api/edge


1.-The bot will remind you to watch your anime on the day is published
2.-You cand add the anime that you to be reminded to you

1.-The list of animes will contain if is being published or has been finished
2.-if the anime has alredy been finished, you can choose a day to be reminded
*/
