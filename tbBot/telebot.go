package tbBot

import (
	"log"
	"time"

	"github.com/DanielUlises98/mytelebot/API"
	"github.com/DanielUlises98/mytelebot/KEYS"
	"github.com/DanielUlises98/mytelebot/models"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	chatID string = ""
	// Universal markup builders.
	inlaneAnime = &tb.ReplyMarkup{}
	animeMenu   = &tb.ReplyMarkup{}

	// Inline buttons.
	//
	// Pressing it will cause the client to
	// send the bot a callback.
	//
	// Make sure Unique stays unique as per button kind,
	// as it has to be for callback routing to work.
	//
	animes = inlaneAnime.Data("Animes", "animes")

	fav      = animeMenu.Data("Favorites", "favorites")
	Status   = animeMenu.Data("Current", "status")
	goBack   = animeMenu.Data("Back Animes", "back")
	addAnime = animeMenu.Data("Add anime", "addanime")
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
	inlaneAnime.Inline(
		inlaneAnime.Row(animes),
	)
	animeMenu.Inline(
		animeMenu.Row(fav, Status),
		animeMenu.Row(addAnime),
		animeMenu.Row(goBack),
	)
	//Simplified way
	bot := TheBot{TB: b, H: API.DBClient{DB: models.InitDB()}}
	//Old way()
	// bot := new(TheBot)
	// bot.TB = b
	// bot.H.DB = models.InitDB()
	//METHODS

	//bot.StartEndPoint()
	bot.TB.Handle("/start", bot.Start)
	bot.TB.Handle("/anime", bot.Anime)
	bot.TB.Handle("/animes", bot.ListOfAnimes) // will display a linearKeyboard of  alredy added by the user

	bot.TB.Handle(&animes, bot.AnimeMenu)
	bot.TB.Handle(&goBack, bot.GoBackButton)
	bot.TB.Handle(&addAnime, bot.AddAnime)

	bot.TB.Handle(tb.OnText, bot.TextFromChat)
	//bot.TB.Handle(tb.OnChosenInlineResult, bot.Anime)

	//bot.QueryKeyboard()
	// Start the bot at the end
	bot.TB.Start()
}
