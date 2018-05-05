package lovepodcast_test

import (
	"github.com/billzhuang/lovepodcast"
	"net/http"
	"os"
	"net/http/httptest"
)

func Entry_httpHandlers() {

	httpHandler := func(w http.ResponseWriter, r *http.Request) {

		cast := lovepodcast.Test()

		// set the Content Type to that of XML
		w.Header().Set("Content-Type", "application/xml")

		// finally, Encode and write the Podcast to the ResponseWriter.
		//
		// a simple pattern is to handle any errors within this check.
		// alternatively if using middleware, you can just return
		// the Podcast entity as it also implements the io.Writer interface
		// that complies with several middleware packages.
		if err := cast.Encode(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	rr := httptest.NewRecorder()
	httpHandler(rr, nil)
	os.Stdout.Write(rr.Body.Bytes())

}
