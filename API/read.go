package API

import (
	"github.com/DanielUlises98/mytelebot/models"
)

func (driver DBClient) UserAnimes(ci string) (animes []models.Anime) {
	user := &models.User{}
	//animes := &[]models.Anime{}
	driver.DB.Where("id = ?", ci).First(&user)
	driver.DB.Model(&user).Association("Animes").Find(&animes)
	return
}
func (driver DBClient) Hours() []models.UserAnimes {
	ua := &[]models.UserAnimes{}
	driver.DB.Where("remind_user = ?", true).Find(&ua)
	return *ua
}
