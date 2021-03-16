package API

import "gopkg.in/tucnak/telebot.v2"

func (driver DBTLClient) Send() {
	spawn := &telebot.User{ID: 695139185}
	driver.TB.Send(spawn, "hola")
}
