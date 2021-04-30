package tbBot

import (
	"fmt"
	"log"
	"time"

	"github.com/DanielUlises98/mytelebot/API"
	"github.com/DanielUlises98/mytelebot/KEYS"
	"github.com/DanielUlises98/mytelebot/models"
	"github.com/DanielUlises98/mytelebot/reminder"
	tb "gopkg.in/tucnak/telebot.v2"
)

func StartBot() {
	b, err := tb.NewBot(tb.Settings{
		Token:  KEYS.TELEBOT_KEY,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot := TheBot{TB: b, H: API.DBClient{DB: models.InitDB()}}

	bot.TB.Handle("/start", bot.Start)
	bot.TB.Handle("/add", bot.SearchResult)
	bot.TB.Handle("/list", bot.AddedList)
	bot.TB.Handle("/cr", bot.ChangeRelease)

	//bot.TB.Handle(&addAnime, bot.AddAnime)

	bot.TB.Handle(tb.OnText, bot.TextFromChat)

	//bot.QueryKeyboard()
	// Start the bot at the end
	reminder.StartReminder(bot.H)
	fmt.Println("Starting Bot")
	bot.TB.Start()

}
