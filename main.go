package main

import (
	"github.com/DanielUlises98/mytelebot/API"
	"github.com/DanielUlises98/mytelebot/models"
	"github.com/DanielUlises98/mytelebot/tbstart"
)

func main() {

	db := models.InitDB()
	bot := tbstart.StartBot()

	caller := API.DBTLClient{DB: db, TB: bot}
	caller.CreateNewUser()
	bot.Start()
}

/*


 */
