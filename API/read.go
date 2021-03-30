package API

import (
	"github.com/DanielUlises98/mytelebot/models"
)

func (driver DBClient) UserAnimes(ci string) (animes []models.Anime) {

	user := &models.User{}
	//animes := &[]models.Anime{}
	driver.DB.Where("chat_id = ?", ci).First(&user)
	driver.DB.Model(&user).Association("Animes").Find(&animes)

	return
}
