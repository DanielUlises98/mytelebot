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

type InnerUA struct {
	UserID     string
	HourRemind string
	WeekDay    string
	Name       string
}

func (driver DBClient) Hours(weekday string) []InnerUA {
	//ua := &[]models.UserAnimes{}
	result := &[]InnerUA{}
	//driver.DB.Where("remind_user = ?", true).Find(&ua)
	driver.DB.Model(&models.UserAnimes{}).
		Select("user_animes.user_id, user_animes.hour_remind, user_animes.week_day, animes.name").
		Where("remind_user = ? AND week_day = ?", true, weekday).
		Joins("left join animes on animes.id = user_animes.anime_id").
		Scan(&result)
	return *result
}
