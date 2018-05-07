package lovepodcast

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"encoding/json"
	"time"
	"io/ioutil"
)

type XimalayaPodcast struct {
	SoundIds string
}

type XimalayaEpisode struct {
	PlayPath64 string `json:"play_path_64"`
	Title string `json:"title"`
	Intro string `json:"intro"`
	CoverUrl142 string `json:"cover_url_142"`

}

func GetPodcast(url string) XimalayaPodcast {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	resultPodcast := XimalayaPodcast{}

	// Find the review items
	doc.Find("div .personal_body").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		soundIds, exists := s.Attr("sound_ids")

		if exists {
			resultPodcast.SoundIds = soundIds
		}

		fmt.Printf("Review %d: %s\n", i, soundIds)
	})

	return resultPodcast
}

func GetEpisode(url string) XimalayaEpisode{
	// Request the HTML page.
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil{
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")


	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	resultEpisode := XimalayaEpisode{}

	err3 := json.Unmarshal(body,  &resultEpisode)

	if err3 != nil{
		log.Fatal(err3)
	}

	return resultEpisode
}
