package API

import (
	"gopkg.in/tucnak/telebot.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (driver DBTLClient) SendMessage(chatId int, message string) {
	spawn := &telebot.User{ID: chatId}
	driver.TB.Send(spawn, message)
}

func (driver DBTLClient) InlineTest() {
	var (
		// Universal markup builders.
		inlanePLate = &tb.ReplyMarkup{}
		inlaneOther = &tb.ReplyMarkup{}
		// Inline buttons.
		//
		// Pressing it will cause the client to
		// send the bot a callback.
		//
		// Make sure Unique stays unique as per button kind,
		// as it has to be for callback routing to work.
		//
		hola  = inlanePLate.Data("hola", "hola")
		adios = inlanePLate.Data("adios", "adios")

		random = inlaneOther.Data("other", "other")
	)
	inlanePLate.Inline(
		inlanePLate.Row(hola, adios),
	)
	inlaneOther.Inline(inlaneOther.Row(random))

	// Command: /start <PAYLOAD>
	driver.TB.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		driver.TB.Send(m.Sender, "Hello!", inlanePLate)
	})

	// On inline button pressed (callback)
	driver.TB.Handle(&hola, func(c *tb.Callback) {
		// ...
		// Always respond!
		driver.TB.Respond(c, &tb.CallbackResponse{ShowAlert: false})

		//EDIT IS THE KEY TO CREATE MENUS
		driver.TB.Edit(c.Message, "Bye!", inlaneOther)
		//driver.TB.EditReplyMarkup(c.Message, inlaneOther)
	})
}
