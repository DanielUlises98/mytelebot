package tbBot

import (
	"log"
	"time"

	"github.com/DanielUlises98/mytelebot/KEYS"
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
	bot := new(TheBot)
	bot.TB = b
	bot.InlineTest()

	// Start the bot at the end
	bot.TB.Start()
}
