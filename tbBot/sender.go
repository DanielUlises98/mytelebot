package tbBot

import (
	"encoding/json"
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (driver TheBot) StartEndPoint() {
	driver.TB.Handle("/start", func(m *tb.Message) {
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
	})
}

func (driver TheBot) SendMessage(chatId int, message string) {
	spawn := &tb.User{ID: chatId}
	driver.TB.Send(spawn, message)
}

// Maybe will work for other things
func (driver TheBot) QueryKeyboard() {
	var (
		inlineKeyboard = &tb.ReplyMarkup{}

		query     = inlineKeyboard.Query("hi", "")
		queryChat = inlineKeyboard.QueryChat("bye", "")
	)

	inlineKeyboard.Inline(
		inlineKeyboard.Row(query, queryChat),
	)

	driver.TB.Handle("/other", func(m *tb.Message) {
		driver.TB.Send(m.Sender, "other", inlineKeyboard)
	})
	driver.TB.Handle(tb.OnQuery, func(q *tb.Query) {
		log.Println(q.Text)
	})
	// driver.TB.Handle(&query, func(q *tb.Query) {
	// 	log.Println(q.Text)
	// })

}

func (driver TheBot) InlineKeyboard() {
	var (
		chatID string = ""
		// Universal markup builders.
		inlaneAnime = &tb.ReplyMarkup{}
		inlaneMenu  = &tb.ReplyMarkup{}

		// Inline buttons.
		//
		// Pressing it will cause the client to
		// send the bot a callback.
		//
		// Make sure Unique stays unique as per button kind,
		// as it has to be for callback routing to work.
		//
		animes = inlaneAnime.Data("Animes", "animes")

		fav      = inlaneMenu.Data("Favorites", "favorites")
		Status   = inlaneMenu.Data("Current", "status")
		goBack   = inlaneMenu.Data("Back Animes", "back")
		addAnime = inlaneMenu.Data("Add anime", "addanime")
	)
	inlaneAnime.Inline(
		inlaneAnime.Row(animes),
	)
	inlaneMenu.Inline(
		inlaneMenu.Row(fav, Status),
		inlaneMenu.Row(addAnime),
		inlaneMenu.Row(goBack),
	)

	// Command: /start <PAYLOAD>
	driver.TB.Handle("/anime", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		driver.TB.Send(m.Sender, "Animes", inlaneAnime)
	})

	// On inline button pressed (callback)
	driver.TB.Handle(&animes, func(c *tb.Callback) {
		// ...
		// Always respond!
		driver.TB.Respond(c, &tb.CallbackResponse{ShowAlert: false})

		//EDIT IS THE KEY TO CREATE MENUS
		driver.TB.Edit(c.Message, "Here is the menu!", inlaneMenu)
		//driver.TB.EditReplyMarkup(c.Message, inlaneOther)
	})

	driver.TB.Handle(&goBack, func(c *tb.Callback) {
		driver.TB.Respond(c, &tb.CallbackResponse{ShowAlert: false})

		driver.TB.Edit(c.Message, "Animes", inlaneAnime)
	})

	driver.TB.Handle(&addAnime, func(c *tb.Callback) {
		driver.TB.Respond(c, &tb.CallbackResponse{ShowAlert: false})
		driver.TB.Send(c.Sender, "Send the name of the anime")

		chatID = ChatID(c.Message.Chat)
	})

	driver.TB.Handle(tb.OnText, func(m *tb.Message) {
		if chatID != "" {
			log.Println("Call the DB", chatID)
			chatID = ""
			return
		}
		log.Println(m.Text)
	})
}

func ChatID(chat *tb.Chat) (ci string) {
	byteC, err := json.Marshal(chat.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	ci = string(byteC)
	return
}
