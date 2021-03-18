package API

import "gopkg.in/tucnak/telebot.v2"

func (driver DBTLClient) SendMessage(chatId int, message string) {
	spawn := &telebot.User{ID: chatId}
	driver.TB.Send(spawn, message)
}
