package kitsu

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

func SearchAnime(animeName string) (animeID string) {
	resp, err := http.Get("https://kitsu.io/api/edge/anime?filter[text]=" + animeName + "&page[limit]=1")
	if err != nil {
		log.Fatal(err, "Couldn't GET the json")
	}
	defer resp.Body.Close()

	wr, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, "Couldn't read the body")
	}
	//log.Println(string(wr))
	anime := &ArrayData{}
	//anime.Data = make([]Anime, 0)

	//anime.Data = &Anime{}
	err_ := json.Unmarshal(wr, anime)
	if err_ != nil {
		log.Fatal(err_, "Couldn't unmarshal the json")
	}

	animeID = anime.Data[0].AnimeID()
	return
}

// var prettyJSON bytes.Buffer
// err_ := json.Indent(&prettyJSON, wr, "", "    ")
// if err_ != nil {
// 	log.Fatal(err_, "Couldn't marshal")
// }
// fmt.Println(string(prettyJSON.Bytes()))
