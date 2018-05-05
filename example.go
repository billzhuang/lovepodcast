package lovepodcast
/*
import (
	"net/http"
	"github.com/eduncan911/podcast"
	"strconv"
	"fmt"
)

// ResponseWriter example using Podcast.Encode(w io.Writer).
//
httpHandler := func(w http.ResponseWriter, r *http.Request) {

	// instantiate a new Podcast
	p := podcast.New(
		"eduncan911 Podcasts",
		"http://eduncan911.com/",
		"An example Podcast",
		&pubDate, &updatedDate,
	)

	// add some channel properties
	p.AddAuthor("Jane Doe", "me@janedoe.com")
	p.AddAtomLink("http://eduncan911.com/feed.rss")
	p.AddImage("http://janedoe.com/i.jpg")
	p.AddSummary(`link <a href="http://example.com">example.com</a>`)
	p.IExplicit = "no"

	for i := int64(1); i < 3; i++ {
		n := strconv.FormatInt(i, 10)
		d := pubDate.AddDate(0, 0, int(i))

		// create an Item
		item := podcast.Item{
			Title:       "Episode " + n,
			Link:        "http://example.com/" + n + ".mp3",
			Description: "Description for Episode " + n,
			PubDate:     &d,
		}
		item.AddImage("http://example.com/episode-" + n + ".png")
		item.AddSummary(`item <a href="http://example.com">example.com</a>`)
		// add a Download to the Item
		item.AddEnclosure("http://e.com/"+n+".mp3", podcast.MP3, 55*(i+1))

		// add the Item and check for validation errors
		if _, err := p.AddItem(item); err != nil {
			fmt.Println(item.Title, ": error", err.Error())
			return
		}
	}

	// set the Content Type to that of XML
	w.Header().Set("Content-Type", "application/xml")

	// finally, Encode and write the Podcast to the ResponseWriter.
	//
	// a simple pattern is to handle any errors within this check.
	// alternatively if using middleware, you can just return
	// the Podcast entity as it also implements the io.Writer interface
	// that complies with several middleware packages.
	if err := p.Encode(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

rr := httptest.NewRecorder()
httpHandler(rr, nil)
os.Stdout.Write(rr.Body.Bytes())
*/