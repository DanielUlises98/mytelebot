package API

import (
	"github.com/DanielUlises98/mytelebot/models"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

// type UserResponse struct {
// 	User models.User //`json:"user"`
// 	Data interface{} `json:"data"`
// }

// CreateNewUser creates a new telegram user
func (driver DBClient) NewUser(username, chatId string) string {

	user := models.User{
		Username: username,
		ID:       chatId,
	}

	userq := &models.User{}
	driver.DB.Where("id = ?", user.ID).First(&userq)
	if !(user.ID == userq.ID) {
		driver.DB.Create(&user)
		return "Welcome"
	} else {
		return "We alredy know each other"
	}
}

//AssociateAnime adds new animes to db and links them to a user
//An associates existing animes to a user
func (driver DBClient) AssociateAnime(ci string, anime models.Anime) bool {
	user := &models.User{}
	driver.DB.Where("id = ?", ci).First(&user)

	animeF := &models.Anime{}
	animesID := []string{anime.ID}

	tx := driver.DB.Where("id = ?", anime.ID).First(&animeF).Error
	if tx != nil {
		driver.DB.Model(&user).Association("Animes").Append(&anime)
		return true
	}

	temp := animeF
	animeF = &models.Anime{}
	driver.DB.Model(&user).Where("id IN ?", animesID).Association("Animes").Find(&animeF)
	if animeF.ID == "" {
		driver.DB.Create(&models.UserAnimes{UserID: user.ID, AnimeID: temp.ID})
		return true
	}
	return false
}

/*
1 anime already exist
2 user wants to add an already existing anime
3
*/

// func (driver DBClient) User(ci string) (user models.User) {
// 	driver.DB.Where("chat_id = ?", ci).First(&user)
// 	return
// }

// func (driver DBClient) Anime(ci string) {
// 	user := driver.User(ci)

// }
