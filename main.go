package main

import "github.com/DanielUlises98/mytelebot/kitsu"

// type Animes interface {
// 	GetType() string
// }
// type Anime struct {
// 	Type string `json:"type"`
// }

// type StartPoint struct {
// 	Data []Anime `json:"data"`
// }

// func (A Anime) GetType() string {
// 	return A.Type
// }
func main() {
	kitsu.Populate()
	// var prettyJSON bytes.Buffer
	// err_ := json.Indent(&prettyJSON, wr, "", "    ")
	// if err_ != nil {
	// 	log.Fatal(err_, "Couldn't marshal")
	// }
	// fmt.Println(string(prettyJSON.Bytes()))
	// db := models.InitDB()
	// bot := tbstart.StartBot()

	// caller := API.DBTLClient{DB: db, TB: bot}
	// caller.CreateNewUser()
	// bot.Start()
}

/*
Make the api out of kitsyu

//BASE API OF KITSU https://kitsu.io/api/edge


1.-The bot will remind you to watch your anime on the day is published
2.-You cand add the anime that you to be reminded to you

1.-The list of animes will contain if is being published or has been finished
2.-if the anime has alredy been finished, you can choose a day to be reminded

*/
