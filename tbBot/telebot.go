package tbBot

import (
	"log"
	"time"

	"github.com/DanielUlises98/mytelebot/API"
	"github.com/DanielUlises98/mytelebot/KEYS"
	"github.com/DanielUlises98/mytelebot/models"
	tb "gopkg.in/tucnak/telebot.v2"
)

func StartBot() {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: "http://195.129.111.17:8012",

		Token:  KEYS.TELEBOT_KEY,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	//Simplified way
	bot := TheBot{TB: b, H: API.DBClient{DB: models.InitDB()}}
	//Old way()
	// bot := new(TheBot)
	// bot.TB = b
	// bot.H.DB = models.InitDB()
	//METHODS
	bot.StartEndPoint()
	bot.InlineKeyboard()
	//bot.QueryKeyboard()
	// Start the bot at the end
	bot.TB.Start()
}
