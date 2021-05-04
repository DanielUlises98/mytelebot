package API

import (
	"testing"

	"github.com/DanielUlises98/mytelebot/KEYS"
	"github.com/DanielUlises98/mytelebot/models"
)

var (
	H DBClient
)

func TestHours(t *testing.T) {
	H.DB = models.InitDB(KEYS.DSN)
	//fmt.Println(H.Hours(), len(H.Hours()))
}
