package lovepodcast

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type XimalayaPodcast struct {
	SoundIds string
}

func New() XimalayaPodcast {
	return XimalayaPodcast{}
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

	resultPodcast := New()

	// Find the review items
	doc.Find(".personal_body").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		soundIds := s.Find("sound_ids").Text()
		resultPodcast.SoundIds = soundIds
		fmt.Printf("Review %d: %s\n", i, soundIds)
	})

	return resultPodcast
}
