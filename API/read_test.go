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
	H.DB = models.InitDB("host=localhost user=duckes password=123456 dbname=mydb port=5432 sslmode=disable")
	fmt.Println(H.Hours(), len(H.Hours()))
}
