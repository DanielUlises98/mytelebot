package API

import (
	"github.com/DanielUlises98/mytelebot/models"
)

func (driver DBClient) UpdateWeekday(idU string, idA string, hr int, weekDay string, remind bool) {
	ua := &models.UserAnimes{}
	driver.DB.Model(&ua).Select("hour_remind", "remind_user", "week_day").Where("user_id = ? AND anime_id = ?", idU, idA).
		Updates(models.UserAnimes{HourRemind: hr, RemindUser: remind, WeekDay: weekDay})
	// driver.DB.Model(&ua).Where("user_id = ? AND anime_id = ?", idU, idA).
	// 	Updates(models.UserAnimes{HourRemind: hr, WeekDay: weekDay})
}
func (driver DBClient) UpdateTz(ui string, tz string) {
	u := &models.User{}
	driver.DB.Where("id = ?", ui).First(&u)
	u.TimeZone = tz
	driver.DB.Save(&u)
}
