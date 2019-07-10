package compose

import (
	"github.com/samcsf/daily-go/pkg/util"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func handlePing(w http.ResponseWriter, req *http.Request) {
	time.Sleep(300 * time.Millisecond)
	w.Write([]byte("Pong"))
}

func duration(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			log.Printf("Request takes %d ms", end.Sub(start).Nanoseconds()/1e6)
		}()
		next(w, req)
		return
	}
}

func addHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Happy", "yo")
		next(w, req)
		return
	}
}

func TestCompose(t *testing.T) {
	ts1 := httptest.NewServer(http.Handler(duration(addHeader(handlePing))))
	defer ts1.Close()
	ts2 := httptest.NewServer(http.Handler(Compose(duration, addHeader)(handlePing)))
	defer ts2.Close()

	resp1, err := http.Get(ts1.URL)
	util.ChkErr(err)

	resp2, err := http.Get(ts2.URL)
	util.ChkErr(err)

	body1, _ := ioutil.ReadAll(resp1.Body)
	body2, _ := ioutil.ReadAll(resp2.Body)
	cmp := strings.Compare(string(body1), string(body2))
	if cmp != 0 {
		log.Println("Body not match")
		t.Fail()
	}

	eql := reflect.DeepEqual(resp1.Header, resp2.Header)
	if !eql {
		log.Printf("Expect: %v, Got: %v", resp1.Header, resp2.Header)
		t.Fail()
	}
}
