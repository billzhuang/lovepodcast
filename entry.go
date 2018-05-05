package lovepodcast

import (
	"fmt"
	"github.com/eduncan911/podcast"
	"time"
)

func NewPodcast(pubDate, updatedDate *time.Time) podcast.Podcast {
	p := podcast.New(
		"eduncan911 Podcasts",
		"http://eduncan911.com/",
		"An example Podcast",
		pubDate, updatedDate,
	)

	// add some channel properties
	p.AddAuthor("Jane Doe", "me@janedoe.com")
	p.AddAtomLink("http://eduncan911.com/feed.rss")
	p.AddImage("http://janedoe.com/i.jpg")
	p.AddSummary(`link <a href="http://example.com">example.com</a>`)
	p.IExplicit = "no"

	return p
}


func NewEpisode(pubDate *time.Time) podcast.Item {
	item := podcast.Item{
		Title:       "Episode ",
		Link:        "http://example.com/.mp3",
		Description: "Description for Episode ",
		PubDate:     pubDate,
	}
	item.AddImage("http://example.com/episode-.png")
	item.AddSummary(`item <a href="http://example.com">example.com</a>`)
	// add a Download to the Item
	item.AddEnclosure("http://e.com/.mp3", podcast.MP3, 55)

	return item
}

func Test() podcast.Podcast{
	now := time.Now()

	cast := NewPodcast(&now, &now)
	episode := NewEpisode(&now)

	// add the Item and check for validation errors
	if _, err := cast.AddItem(episode); err != nil {
		fmt.Println(episode.Title, ": error", err.Error())
		//return
	}

	return cast

	/*if err := cast.Encode(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}*/
}
