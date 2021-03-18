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
	IdAnime string     `json:"id"`
	Attri   Attributes `json:"attributes"`
}

type ArrayData struct {
	Data []Anime `json:"data"`
}
type OData struct {
	Data Anime `json:"data"`
}

type Animes interface {
	GetAnimeId() int
	GetAttributes() Attributes
}

func (A Anime) GetAnimeId() string {
	return A.IdAnime
}
func (A Anime) GetAttributes() Attributes {
	return A.Attri
}

func SearchAnime(animeName string) {
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

	log.Println(anime.Data[0].GetAnimeId())
	log.Println(anime.Data[0].GetAttributes().Titles.English)
}

// var prettyJSON bytes.Buffer
// err_ := json.Indent(&prettyJSON, wr, "", "    ")
// if err_ != nil {
// 	log.Fatal(err_, "Couldn't marshal")
// }
// fmt.Println(string(prettyJSON.Bytes()))
