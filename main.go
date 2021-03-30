package main

import (
	"github.com/DanielUlises98/mytelebot/kitsu"
)

// var (
// 	api API.DBClient
// )

func main() {
	//tbBot.StartBot()
	//api = API.DBClient{DB: models.InitDB()}
	kitsu.SearchAnime("kimetsu no yaiba")
	// bears(&Bear{})
	// bears(&PolarBear{})
	// t, err := time.Parse("2006-01-02", "2019-04-06")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// fmt.Printf("%+v\n", t)
	// fmt.Printf("%+v\n", time.Now())
}

/*
Make the api out of kitsyu

//BASE API OF KITSU https://kitsu.io/api/edge


1.-The bot will remind you to watch your anime on the day is published
2.-You cand add the anime that you to be reminded to you

1.-The list of animes will contain if is being published or has been finished
2.-if the anime has alredy been finished, you can choose a day to be reminded

*/
