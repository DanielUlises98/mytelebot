package main

import "github.com/DanielUlises98/mytelebot/tbBot"

// var (
// 	api API.DBClient
// )

func main() {
	tbBot.StartBot()
	//t, _ := time.Parse(time.Kitchen, "1:46PM")

	// t2 := time.Now().Add(time.Hour * 24)
	// fmt.Println(t2.Clock())
	// fmt.Println(time.Until(t2))
	// timer := time.NewTimer(time.Until(t2))
	// <-timer.C
	// fmt.Println("Timer triggered")

	//ticker := time.NewTicker(5 * time.Second)

}

/*
Make the api out of kitsyu

//BASE API OF KITSU https://kitsu.io/api/edge


1.-The bot will remind you to watch your anime on the day is published
2.-You cand add the anime that you to be reminded to you

1.-The list of animes will contain if is being published or has been finished
2.-if the anime has alredy been finished, you can choose a day to be reminded
*/
