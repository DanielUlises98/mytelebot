package reminder

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DanielUlises98/mytelebot/tbBot"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

type TimeChan struct {
	t *time.Timer
}

var (
	driver     tbBot.TheBot
	daysOfWeek = map[string]time.Weekday{}
	logger     *log.Logger
)

func init() {
	for d := time.Sunday; d <= time.Saturday; d++ {
		daysOfWeek[d.String()] = d
	}
}
func InitVars(db *gorm.DB, bot *tb.Bot) {
	driver.H.DB = db
	driver.TB = bot
	logger = log.New(os.Stdout, "", 0)
	//bDB := BotDB{bot: tbBot.TheBot{TB: bot, H: API.DBClient{DB: db}}}
	go startReminder()
}

func getNextDuration() time.Duration {
	nowUtc := time.Now().UTC()
	next := time.Date(nowUtc.Year(), nowUtc.Month(), nowUtc.Day(), nowUtc.Hour()+1, 00, 00, 0, nowUtc.Location())
	// if next.Before(nowUtc) {
	// 	next = next.Add(INTERVAL)
	// }
	logger.Printf("Current utc time %s \n", nowUtc)
	//next := nowUtc.Add(time.Minute * 5)
	logger.Printf("It is going to start %s \n", next)
	log.Println(next.Sub(nowUtc))
	//dura :=
	return next.Sub(nowUtc)
}

func newJobTimer() TimeChan {
	logger.Printf("Starting the Timer for the reminders\n")
	return TimeChan{time.NewTimer(getNextDuration())}
}
func (j TimeChan) updateJobTimer() {
	j.t.Reset(getNextDuration())
}

func startReminder() {
	jt := newJobTimer()
	for {
		<-jt.t.C
		logger.Printf("%s - JUST TICKED\n", time.Now().UTC())
		listOfReminds()
		jt.updateJobTimer()
	}
}

func listOfReminds() {
	utz := driver.H.Hours()
	for _, userTime := range utz {
		load, err := time.LoadLocation(userTime.TimeZone)
		if err != nil {
			log.Printf("%s\n", err)
		}

		tz := time.Now().In(load)
		tz = time.Date(tz.Year(), tz.Month(), tz.Day(), userTime.HourRemind, tz.Minute(), tz.Second(), tz.Nanosecond(), load)
		weekday, err := parseWeekday(userTime.WeekDay)

		if err != nil {
			log.Printf("%s", err)
		}

		tz = trimSec(tz)
		utc := trimSec(time.Now().UTC())
		if weekday == tz.Weekday() {
			log.Printf("Time zone %s is equal to utc %s", tz, utc)
			if tz.Equal(utc) {
				log.Printf("Reminding user %s for anime %s", userTime.UserID, userTime.Name)
				driver.SendUser(userTime.UserID, userTime.Name, tz.Format("Monday 15:04"))
			}
		}
	}
}

func trimSec(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}
func parseWeekday(wd string) (time.Weekday, error) {
	if weekday, ok := daysOfWeek[wd]; ok {
		return weekday, nil
	}
	return time.Sunday, fmt.Errorf("invalid weekday %s", wd)
}
