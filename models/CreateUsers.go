package models

// type UserResponse struct {
// 	User models.User //`json:"user"`
// 	Data interface{} `json:"data"`
// }

// CreateNewUser creates a new telegram user
func (driver DBTLClient) CreateNewUser(username, chatId string) {

	user := User{
		Username: string(username),
		ChatId:   string(chatId),
	}

	userq := &User{}
	driver.DB.Where("chat_id = ?", user.ChatId).First(&userq)
	if !(user.ChatId == userq.ChatId) {
		driver.DB.Create(&user)
	}
}
