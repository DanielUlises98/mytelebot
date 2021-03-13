package telebot

import (
	"encoding/json"
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
		return
	}
	// !A VERY IMPORTANT HINT here := tb.User{ID: }

	//sendme:=tb.Message{Chat: }

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		a, err := json.Marshal(m.Chat.ID)
		if err != nil {
			log.Fatal(err, "JSON marshal failed")
		}
		log.Println(string(a))
		b.Send(m.Sender, string(a))
	})

	b.Start()
}
