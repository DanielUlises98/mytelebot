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
	Slug         string        `json:"slug"`
	Titles       []Titles      `json:"titles"`
	StartDate    string        `json:"startDate"`
	EndDate      string        `json:"endDate"`
	Status       string        `json:"status"`
	PosterImage  []PosterImage `json:"posterImage"`
	EpisodeCount int           `json:"episodeCount"`
	ShowType     string        `json:"showType"`
	NSFW         bool          `json:"nsfw"`
}

type Anime struct {
	IdAnime    string     `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type MainData struct {
	Data *Anime `json:"data"`
}

type Animes interface {
	GetAnimeId() int
	GetAttributes() Attributes
}

func (A Anime) GetAnimeId() string {
	return A.IdAnime
}

func Populate() {
	resp, err := http.Get("https://kitsu.io/api/edge/anime/41370")
	if err != nil {
		log.Fatal(err, "get failed")
	}
	defer resp.Body.Close()
	wr, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, "Couldn't read the body")
	}
	//log.Println(string(wr))
	anime := &MainData{}
	//anime.Data = make([]Anime, 0)

	anime.Data = &Anime{}
	err_ := json.Unmarshal(wr, anime)
	if err_ != nil {
		log.Fatal(err_)
	}
	log.Println(anime.Data.GetAnimeId())
}
