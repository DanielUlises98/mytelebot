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

type UserTZData struct {
	UserID     string
	HourRemind string
	WeekDay    string
	TimeZone   string
	Name       string
}

func (driver DBClient) Hours(weekday string) []UserTZData {
	//ua := &[]models.UserAnimes{}
	result := &[]UserTZData{}
	//driver.DB.Where("remind_user = ?", true).Find(&ua)
	driver.DB.Model(&models.UserAnimes{}).
		Select("user_animes.user_id, user_animes.hour_remind, user_animes.week_day, user.time_zone, animes.name").
		Where("remind_user = ? AND week_day = ?", true, weekday).
		Joins("left join animes on animes.id = user_animes.anime_id").
		Scan(&result)
	return *result
}
func (driver DBClient) UserTz(ui string) (bool, string) {
	u := &models.User{}
	driver.DB.Where("id = ?", ui).First(&u)
	if u.TimeZone == "" {
		return false, ""
	}
	return true, u.TimeZone
}
