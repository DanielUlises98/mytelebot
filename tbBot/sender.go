package tbBot

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/DanielUlises98/mytelebot/kitsu"
	tb "gopkg.in/tucnak/telebot.v2"
)

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

func (driver TheBot) Animes(m *tb.Message) {
	inlaneAnimes := &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	animes := driver.H.UserAnimes(ChatID(m.Chat))

	animesRows := make([]tb.Row, len(animes))

	for i, anime := range animes {
		animesRows[i] = tb.Row{inlaneAnimes.Data(anime.IdAnime, fmt.Sprint(i))}
	}

	inlaneAnimes.Inline(animesRows...)
	driver.TB.Send(m.Sender, "Animes", inlaneAnimes)
}

func (driver TheBot) Anime(m *tb.Message) {
	if !m.Private() {
		return
	}
	driver.TB.Send(m.Sender, "Animes", inlaneAnime)
}

func (driver TheBot) AnimeButtons(c *tb.Callback) {
	// ...
	// Always respond!
	driver.TB.Respond(c, &tb.CallbackResponse{ShowAlert: false})

	//EDIT IS THE KEY TO CREATE MENUS
	driver.TB.Edit(c.Message, "Here is the menu!", inlaneMenu)
	//driver.TB.EditReplyMarkup(c.Message, inlaneOther)
}

func (driver TheBot) GoBackButton(c *tb.Callback) {
	driver.TB.Respond(c, &tb.CallbackResponse{ShowAlert: false})

	driver.TB.Edit(c.Message, "Animes", inlaneAnime)
}

func (driver TheBot) AddAnime(c *tb.Callback) {
	driver.TB.Respond(c, &tb.CallbackResponse{ShowAlert: false})
	driver.TB.Send(c.Sender, "Send the name of the anime")

	chatID = ChatID(c.Message.Chat)
}

func (driver TheBot) TextFromChat(m *tb.Message) {
	if chatID != "" {
		//animeID := models.Anime{IdAnime: kitsu.SearchAnime(m.Text)}
		driver.H.AssociateAnime(chatID, kitsu.SearchAnime(m.Text))
		chatID = ""
		return
	}
}

func ChatID(chat *tb.Chat) (ci string) {
	byteC, err := json.Marshal(chat.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	ci = string(byteC)
	return
}

func (driver TheBot) SendUser(chatId int, message string) {
	user := &tb.User{ID: chatId}
	driver.TB.Send(user, message)
}

////////////////////////

// // Maybe will work for other things
// func (driver TheBot) QueryKeyboard() {
// 	var (
// 		inlineKeyboard = &tb.ReplyMarkup{}

// 		query     = inlineKeyboard.Query("hi", "")
// 		queryChat = inlineKeyboard.QueryChat("bye", "")
// 	)

// 	inlineKeyboard.Inline(
// 		inlineKeyboard.Row(query, queryChat),
// 	)

// 	driver.TB.Handle("/other", func(m *tb.Message) {
// 		driver.TB.Send(m.Sender, "other", inlineKeyboard)
// 	})
// 	driver.TB.Handle(tb.OnQuery, func(q *tb.Query) {
// 		log.Println(q.Text)
// 	})
// }
