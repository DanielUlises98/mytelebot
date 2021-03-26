package API

import (
	"log"

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
		ChatId:   chatId,
	}

	log.Println(driver.DB, "HI THERE")
	userq := &models.User{}
	driver.DB.Where("chat_id = ?", user.ChatId).First(&userq)
	if !(user.ChatId == userq.ChatId) {
		driver.DB.Create(&user)
		return "Welcome"
	} else {
		return "We alredy know each other"
	}
}

func (driver DBClient) AssociateAnime(ci string, anime models.Anime) {
	user := &models.User{}
	driver.DB.Where("chat_id = ?", ci).First(&user)

	animeF := &models.Anime{}
	animesID := []string{anime.IdAnime}

	tx := driver.DB.Where(anime).First(&animeF)
	if tx.Error != nil {
		driver.DB.Model(&user).Association("Animes").Append(&anime)
		return
	}
	temp := animeF
	animeF = &models.Anime{}
	driver.DB.Model(&user).Where("id_anime IN ?", animesID).Association("Animes").Find(&animeF)
	if animeF.IdAnime == "" {
		//driver.DB.Model(&user).Association("Animes").Append(&temp)
		//driver.DB.Model(&temp).Association("Users").Append(&user)
		driver.DB.Create(&models.UserAnimes{UserID: user.ID, AnimeID: temp.ID})
		return
	}
	//fmt.Printf("%+v\n", animeF)
	// animeF = &models.Anime{}
	// driver.DB.Model(&user).Where("id_anime IN ?", animesID).Association("Animes").Find(&animeF)
	// driver.DB.Model(&user).Association("Animes").Append([]models.Anime{{IdAnime: "202022"}, {IdAnime: "52323"}})
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
