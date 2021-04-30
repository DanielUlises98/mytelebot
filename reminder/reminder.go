package reminder

import (
	"fmt"
	"sync"
	"time"

	"github.com/DanielUlises98/mytelebot/API"
	tbbot "github.com/DanielUlises98/mytelebot/tbBot"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

const INTERVAL = time.Hour * 24
const HOUR_TO_FETCH = 23
const MIN_TO_FETCH = 59
const SEC_TO_FETCH = 59

type TimeChan struct {
	t *time.Timer
}
type BotDB struct {
	bot tbbot.TheBot
}

var (
	wg       sync.WaitGroup
	lenTimes int
	ua       []API.InnerUA
)

func (driver BotDB) setUpReminder() {
	ua = driver.bot.H.Hours()
	lenTimes = len(ua)
	driver.setWorkers(setTimers(gatherHr(ua)), ua)
}

func gatherHr(ua []API.InnerUA) []time.Time {
	ts := make([]time.Time, lenTimes)
	now := time.Now()
	for i, item := range ua {
		t, err := time.Parse(time.Kitchen, item.HourRemind)
		if err != nil {
			fmt.Println(err)
		}
		t = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location())
		ts[i] = t
	}
	return ts
	//fmt.Println(time.Until(t))
}

func setTimers(t []time.Time) []TimeChan {
	tr := make([]TimeChan, lenTimes)
	for i := range t {
		tr[i].t = time.NewTimer(time.Until(t[i]))
	}
	return tr
}

func (driver BotDB) setWorkers(tr []TimeChan, ua []API.InnerUA) {
	wg.Add(lenTimes)
	for i := range tr {
		go func(t *time.Timer, userId, name string) {
			<-t.C
			driver.bot.SendUser(userId, name)
			wg.Done()
		}(tr[i].t, ua[i].UserID, ua[i].Name)
	}
	wg.Wait()
}

func getNextDuration() time.Duration {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_FETCH, MIN_TO_FETCH, SEC_TO_FETCH, 0, now.Location())
	if next.Before(now) {
		next = next.Add(INTERVAL)
	}
	fmt.Println(time.Until(next), " Is going to tick")
	return time.Until(next)
}
func newJobTimer() TimeChan {
	fmt.Println("New timer")
	return TimeChan{time.NewTimer(getNextDuration())}
}
func (j TimeChan) updateJobTimer() {
	j.t.Reset(getNextDuration())
}

func (driver BotDB) StartReminder() {
	jt := newJobTimer()
	for {
		<-jt.t.C
		fmt.Println(time.Now(), "- JUST TICKED")
		driver.setUpReminder()
		jt.updateJobTimer()
	}
}
func Init(db *gorm.DB, bot *tb.Bot) {
	bDB := BotDB{bot: tbbot.TheBot{TB: bot, H: API.DBClient{DB: db}}}
	bDB.StartReminder()
}
