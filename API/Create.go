package API

import (
	"encoding/json"
	"log"

	"github.com/DanielUlises98/mytelebot/models"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

type DBTLClient struct {
	DB *gorm.DB
	TB *tb.Bot
}

// type UserResponse struct {
// 	User models.User //`json:"user"`
// 	Data interface{} `json:"data"`
// }

// CreateNewUser creates a new telegram user
func (driver DBTLClient) CreateNewUser() {

	driver.TB.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		chatId, err := json.Marshal(m.Chat.ID)
		if err != nil {
			log.Fatal(err, "there is something wrong with marshal of chatid ")
		}
		username, err := json.Marshal(m.Chat.Username)
		if err != nil {
			log.Fatal(err, "there is something wrong with marshal of username ")
		}
		user := models.User{
			Username: string(username),
			ChatId:   string(chatId),
		}

		userq := &models.User{}
		driver.DB.Where("chat_id = ?", user.ChatId).First(&userq)
		if !(user.ChatId == userq.ChatId) {
			driver.DB.Create(&user)
			driver.TB.Send(m.Sender, "Hi newcomer!")
		} else {
			driver.TB.Send(m.Sender, "You are alredy subscribed")
		}
	})
}
