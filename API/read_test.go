package API

import (
	"fmt"
	"testing"

	"github.com/DanielUlises98/mytelebot/models"
)

var (
	H DBClient
)

func TestHours(t *testing.T) {
	H.DB = models.InitDB()
	fmt.Println(H.Hours(), len(H.Hours()))
}
