package kitsu

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/DanielUlises98/mytelebot/models"
)

type Titles struct {
	English string `json:"en"`
	EnJp    string `json:"en_jp"`
}

type PosterImage struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

type Attributes struct {
	Slug         string      `json:"slug"`
	Titles       Titles      `json:"titles"`
	StartDate    string      `json:"startDate"`
	EndDate      string      `json:"endDate"`
	Status       string      `json:"status"`
	PosterImage  PosterImage `json:"posterImage"`
	EpisodeCount int         `json:"episodeCount"`
	ShowType     string      `json:"showType"`
	NSFW         bool        `json:"nsfw"`
}

type Anime struct {
	IDAnime string     `json:"id"`
	Attri   Attributes `json:"attributes"`
}

type ArrayData struct {
	Data []Anime `json:"data"`
}
type OData struct {
	Data Anime `json:"data"`
}
type DataParent struct {
	Array ArrayData
	One   OData
}
type Animes interface {
	AnimeID() int
	Attributes() Attributes
}

func (A Anime) AnimeID() string {
	return A.IDAnime
}
func (A Anime) Attributes() Attributes {
	return A.Attri
}

func SearchAnime(animeName string) (anime []models.Anime) {

	animeName = strings.Replace(animeName, " ", "%20", -1)

	// I HAVE TO PUT BY HAND SPECIAL CHARACTERS IN ASCII code
	resp, err := http.Get("https://kitsu.io/api/edge/anime?filter[subtype]=TV&filter[text]=" + animeName)
	if err != nil {
		log.Fatal(err, "Couldn't reach the KITSU API")
	}
	defer resp.Body.Close()
	dt := &ArrayData{}
	jsonUnmarshal(resp.Body, dt)
	return animeWrapper(dt)
}

func SearchAnimeByID(animeID string) {
	resp, err := http.Get("https://kitsu.io/api/edge/anime/" + animeID)
	if err != nil {
		log.Fatal(err, " Couldn't GET the json")
	}
	defer resp.Body.Close()
	anime := &OData{}
	jsonUnmarshal(resp.Body, anime)
	//fmt.Printf("%+v\n", anime.Data.Attributes().Titles)
}

func jsonUnmarshal(r io.Reader, anime interface{}) {
	wr, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err, " Couldn't read the body")
	}

	// DEBUG wr If somethings not working
	//fmt.Printf("JSON BODY %+v\n", string(wr))

	err = json.Unmarshal(wr, &anime)
	if err != nil {
		log.Printf("verbose error info: %#v", err)
		log.Fatal(err.Error(), " Couldn't unmarshal the json")
	}
}

const (
	startDefaultEp = 1
)

func animeWrapper(dt *ArrayData) (anime []models.Anime) {
	anime = make([]models.Anime, len(dt.Data))
	for i, v := range dt.Data {
		anime[i] = models.Anime{
			Episodes:       uint(v.Attributes().EpisodeCount),
			ID:             v.AnimeID(),
			Name:           v.Attributes().Titles.EnJp,
			ImageMedium:    v.Attributes().PosterImage.Medium,
			ImageOriginal:  v.Attributes().PosterImage.Original,
			Status:         isStatus(v.Attributes().Status),
			StartDate:      v.Attributes().StartDate,
			EndDate:        v.Attributes().EndDate,
			CurrentEpisode: startDefaultEp,
			CreatedAt:      time.Now(),
		}
	}
	return
}

func isStatus(status string) bool {
	return status == "current"
}

// var prettyJSON bytes.Buffer
// err_ := json.Indent(&prettyJSON, wr, "", "    ")
// if err_ != nil {
// 	log.Fatal(err_, "Couldn't marshal")
// }
// fmt.Println(string(prettyJSON.Bytes()))
