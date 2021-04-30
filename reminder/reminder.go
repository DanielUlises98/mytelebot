package reminder

import (
	"fmt"
	"sync"
	"time"

	"github.com/DanielUlises98/mytelebot/API"
	"github.com/DanielUlises98/mytelebot/models"
)

const HOUR_TO_FETCH = 22
const MIN_TO_FETCH = 30
const SEC_TO_FETCH = 03

type tickerJob struct {
	tj *time.Timer
}
type TimeChan struct {
	t *time.Timer
}

var (
	wg       sync.WaitGroup
	lenTimes int
)

func setUpReminder(remind API.DBClient) {
	ua := remind.Hours()
	lenTimes = len(ua)
	setWorkers(setTimers(gatherHr(ua)))
}

func gatherHr(ua []models.UserAnimes) []time.Time {
	ts := make([]time.Time, lenTimes)
	for i, item := range ua {
		t, err := time.Parse(time.Kitchen, item.HourRemind)
		if err != nil {
			fmt.Println(err)
		}
		t = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), t.Hour(), t.Minute(), 0, 0, time.Now().Location())
		ts[i] = t
	}
	return ts
	//fmt.Println(time.Until(t))
}

func setTimers(t []time.Time) []TimeChan {
	tr := make([]TimeChan, 1)
	for i := range t {
		tr[i].t = time.NewTimer(time.Until(t[i]))
	}
	return tr
}

func setWorkers(tr []TimeChan) {
	wg.Add(lenTimes)
	for i := range tr {
		go func(t *time.Timer) {
			<-t.C
			fmt.Println("end of goroutine")
			wg.Done()
		}(tr[i].t)
	}
	wg.Wait()
}

func getNextDuration() time.Duration {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_FETCH, MIN_TO_FETCH, SEC_TO_FETCH, 0, now.Location())
	if next.Before(now) {
		next = next.Add(time.Hour * 24)
	}
	fmt.Println(time.Until(next), " Is going to tick")
	return time.Until(next)
}
func newJobTimer() tickerJob {
	fmt.Println("New timer")
	return tickerJob{time.NewTimer(getNextDuration())}
}
func (j tickerJob) updateJobTimer() {
	j.tj.Reset(getNextDuration())
}

func StartReminder(db API.DBClient) {
	jt := newJobTimer()
	for {
		<-jt.tj.C
		fmt.Println(time.Now(), "- JUST TICKED")
		setUpReminder(db)
		jt.updateJobTimer()
	}
}
