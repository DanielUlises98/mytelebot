package reminder

import (
	"log"
	"os"
	"time"

	"github.com/DanielUlises98/mytelebot/API"
	"github.com/DanielUlises98/mytelebot/tbBot"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

const INTERVAL = time.Hour * 24
const HOUR_TO_FETCH = 03
const MIN_TO_FETCH = 26
const SEC_TO_FETCH = 00

type TimeChan struct {
	t *time.Timer
}
type BotDB struct {
	bot tbBot.TheBot
}
type timesZones struct {
	time     time.Time
	timeZone time.Time
}

var (
	//wg       sync.WaitGroup
	lenTimes int
	ua       []API.UserTZData
	logger   *log.Logger
)

func (driver BotDB) setUpReminder() {
	logger.Printf("Initiating the reminder proces\n")
	ua = driver.bot.H.Hours()
	lenTimes = len(ua)
	driver.setWorkers(setTimers(gatherHr(ua)), ua)
}

func gatherHr(ua []API.UserTZData) []timesZones {
	//ts := make([]time.Time, lenTimes)
	tz := make([]timesZones, lenTimes)
	for i, item := range ua {
		load, err := time.LoadLocation(item.TimeZone)
		if err != nil {
			log.Println(err, "error when loading a location")
		}
		nowCurrent := time.Now().In(load)
		//now := time.Now()
		t, err := time.Parse(time.Kitchen, item.HourRemind)
		if err != nil {
			logger.Println(err)
		}
		t = time.Date(nowCurrent.Year(), nowCurrent.Month(), nowCurrent.Day(), t.Hour(), t.Minute(), 0, 0, load)
		if t.Sub(nowCurrent) < time.Duration(time.Hour*24) && nowCurrent.Weekday().String() == item.WeekDay {
			tz[i].time = t
			tz[i].timeZone = nowCurrent
		}
	}
	return tz
	//logger.Println(time.Until(t))
}

func setTimers(tz []timesZones) []TimeChan {
	logger.Printf("Setting timers for the reminders\n")
	tr := make([]TimeChan, lenTimes)
	for i := range tz {
		//tr[i].t = time.NewTimer(time.Until(t[i]))
		tr[i].t = time.NewTimer(tz[i].time.Sub(tz[i].timeZone))
	}
	return tr
}

func (driver BotDB) setWorkers(tr []TimeChan, ua []API.UserTZData) {
	//wg.Add(lenTimes)
	logger.Printf("Setting reminders\n")
	for i := range tr {
		go func(t *time.Timer, userId, name string) {
			logger.Printf("Will remind %s the anime %s\n", userId, name)
			<-t.C
			driver.bot.SendUser(userId, name)
			logger.Printf("%s Reminded", userId)
			//		wg.Done()
		}(tr[i].t, ua[i].UserID, ua[i].Name)
	}
	//wg.Wait()
}

func getNextDuration() time.Duration {
	nowUtc := time.Now().UTC()
	next := time.Date(nowUtc.Year(), nowUtc.Month(), nowUtc.Day(), HOUR_TO_FETCH, MIN_TO_FETCH, SEC_TO_FETCH, 0, nowUtc.Location())
	if next.Before(nowUtc) {
		next = next.Add(INTERVAL)
	}
	logger.Printf("Current utc time %s \n", nowUtc)
	logger.Printf("%s It is going to start the timers \n", next.Sub(nowUtc))
	return next.Sub(nowUtc)
}
func newJobTimer() TimeChan {
	logger.Printf("Starting the Timer for the reminders\n")
	return TimeChan{time.NewTimer(getNextDuration())}
}
func (j TimeChan) updateJobTimer() {
	j.t.Reset(getNextDuration())
}

func (driver BotDB) StartReminder() {
	jt := newJobTimer()
	for {
		<-jt.t.C
		logger.Printf("%s - JUST TICKED\n", time.Now().UTC())
		driver.setUpReminder()
		jt.updateJobTimer()
	}
}
func Init(db *gorm.DB, bot *tb.Bot) {
	logger = log.New(os.Stdout, "", 0)
	bDB := BotDB{bot: tbBot.TheBot{TB: bot, H: API.DBClient{DB: db}}}
	go bDB.StartReminder()
}
