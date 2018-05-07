package lovepodcast_test

import (
	"net/http"
	"os"
	"net/http/httptest"
	"testing"
	"github.com/billzhuang/lovepodcast"
)

func Entry_httpHandlers(t *testing.T) {

	httpHandler := func(w http.ResponseWriter, r *http.Request) {

		cast :=  lovepodcast.Test()

		// set the Content Type to that of XML
		w.Header().Set("Content-Type", "application/xml")

		// finally, Encode and write the Podcast to the ResponseWriter.
		//
		// a simple pattern is to handle any errors within this check.
		// alternatively if using middleware, you can just return
		// the Podcast entity as it also implements the io.Writer interface
		// that complies with several middleware packages.
		if err := cast.Encode(w); err != nil {
			t.Error("error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	rr := httptest.NewRecorder()
	httpHandler(rr, nil)
	os.Stdout.Write(rr.Body.Bytes())

}
