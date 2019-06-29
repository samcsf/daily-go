package compose

import (
	"github.com/samcsf/daily-go/pkg/util"
	// 	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func handlePing(w http.ResponseWriter, req *http.Request) {
	time.Sleep(500 * time.Millisecond)
	w.Write([]byte("Pong"))
}

func duration(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now().Nanosecond() / 1000000
		defer func() {
			log.Printf("Request takes %d ms", time.Now().Nanosecond()/1000000-start)
		}()
		next(w, req)
		return
	}
}

func filter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Filtered!")
		next(w, req)
		return
	}
}

func TestCompose(t *testing.T) {
	// ts := httptest.NewServer(http.Handler(duration(filter(handlePing))))
	ts := httptest.NewServer(http.Handler(Compose(duration, filter, handlePing)))
	defer ts.Close()

	_, err := http.Get(ts.URL)
	util.ChkErr(err)

}
