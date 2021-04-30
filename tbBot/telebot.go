package tbBot

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/DanielUlises98/mytelebot/API"
	"github.com/DanielUlises98/mytelebot/KEYS"
	"github.com/DanielUlises98/mytelebot/kitsu"
	"github.com/DanielUlises98/mytelebot/models"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

var (
	ma  []models.Anime
	cid string = ""
)

type TheBot struct {
	TB *tb.Bot
	H  API.DBClient
}

func StartBot() *tb.Bot {
	b, err := tb.NewBot(tb.Settings{
		Token:  KEYS.TELEBOT_KEY,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func InitHandlers(db *gorm.DB, b *tb.Bot) {
	bot := TheBot{TB: b, H: API.DBClient{DB: db}}

	bot.TB.Handle("/start", bot.Start)
	bot.TB.Handle("/add", bot.SearchResult)
	bot.TB.Handle("/list", bot.AddedList)
	bot.TB.Handle("/cr", bot.ChangeRelease)

	//bot.TB.Handle(&addAnime, bot.AddAnime)

	bot.TB.Handle(tb.OnText, bot.TextFromChat)

	//bot.QueryKeyboard()
	// Start the bot at the end
	fmt.Println("Starting Bot")
}

func (driver TheBot) Start(m *tb.Message) {
	if !m.Private() {
		return
	}

	username, err := json.Marshal(m.Chat.Username)
	if err != nil {
		log.Fatal(err, "there is something wrong with marshal of username ")
	}

	chatId, err := json.Marshal(m.Chat.ID)
	if err != nil {
		log.Fatal(err, "there is something wrong with marshal of chatid ")
	}

	result := driver.H.NewUser(string(username), string(chatId))
	driver.TB.Send(m.Sender, result)
}

func (driver TheBot) SearchResult(m *tb.Message) {
	if m.Payload != "" {
		inlaneAnimes := &tb.ReplyMarkup{
			ResizeReplyKeyboard: true,
			ForceReply:          true,
			OneTimeKeyboard:     true,
			ReplyKeyboardRemove: true,
		}
		//animes := driver.H.UserAnimes(ChatID(m.Chat))
		ma = kitsu.SearchAnime(m.Payload)

		animesRows := make([]tb.Row, len(ma))

		for i, anime := range ma {
			animesRows[i] = tb.Row{inlaneAnimes.Text(anime.Name)}
		}
		cid = chatID(m.Chat)
		inlaneAnimes.Reply(animesRows...)
		driver.TB.Send(m.Sender, "Animes", inlaneAnimes)
	}
	//driver.TB.Send(m.Sender, "test", replyer)
}

func (driver TheBot) AddedList(m *tb.Message) {
	myList := driver.H.UserAnimes(chatID(m.Chat))
	var animes string
	for _, item := range myList {
		animes += "ID : [" + item.ID + "] " + item.Name + "\n"
	}
	driver.TB.Send(m.Sender, animes)
}

//command -ID -hour -weekday -remind (number or text)
func (driver TheBot) ChangeRelease(m *tb.Message) {
	remind := false
	idDay := strings.Split(m.Payload, " ")
	if len(idDay) != 4 {
		driver.TB.Send(m.Sender, "The information is incomplete, I can't procede with the update")
		return
	}
	wd, err := strconv.ParseInt(idDay[2], 10, 8)
	if err != nil {
		log.Println(err, " Value out of range")
		return
	}
	if idDay[3] == "T" {
		remind = true
	}
	t, err := time.Parse(time.Kitchen, idDay[1])
	if err != nil {
		fmt.Printf("%+v\n", err)
		driver.TB.Send(m.Sender, "Please reachek the hour , shold be typed in this format00:00PM/AM")
	} else if wd <= 0 || wd >= 8 {
		driver.TB.Send(m.Sender, "Please reachek the day, choose in a range of 1 to 7 1 being monday")
	} else {
		driver.H.UpdateWeekday(chatID(m.Chat), idDay[0], t.Format(time.Kitchen), int8(wd), remind)
	}
}

func (driver TheBot) TextFromChat(m *tb.Message) {
	if cid != "" {
		fmt.Println(m.Text)
		if name, ok := driver.findAnime(m.Text); ok {
			driver.TB.Send(m.Sender, name+" was succesfully added")
			cid = ""
			return
		} else {
			driver.TB.Send(m.Sender, "Couldn't add  "+name)
			cid = ""
		}
		//animeID := models.Anime{IdAnime: kitsu.SearchAnime(m.Text)}
	}
}

func (driver TheBot) findAnime(message string) (name string, found bool) {
	for _, anime := range ma {
		if anime.Name == message {
			name = anime.Name
			found = driver.H.AssociateAnime(cid, anime)
			return
		}
		name = anime.Name
	}
	return
}

func (driver TheBot) SendUser(userid string, name string) {
	ui, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user := &tb.Chat{ID: ui}
	driver.TB.Send(user, name)
}
func chatID(chat *tb.Chat) (ci string) {
	byteC, err := json.Marshal(chat.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	ci = string(byteC)
	return
}
