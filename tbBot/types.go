package tbBot

import (
	"github.com/DanielUlises98/mytelebot/API"
	tb "gopkg.in/tucnak/telebot.v2"
)

type TheBot struct {
	TB *tb.Bot
	H  API.DBClient
}
