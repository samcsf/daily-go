package controller

import (
	"github.com/samcsf/daily-go/pkg/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var log = logrus.New()

func init() {
	file, err := os.OpenFile("blog_logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	util.ChkErr(err)

	log.Out = file
}

type Middlewares struct{}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (mdw *Middlewares) Duration(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// hack to save the status code
		rec := &statusRecorder{w, 200}
		next(rec, r)

		end := time.Now()
		log.Printf("| %dms | %d | %s %s ", end.Sub(start).Nanoseconds()/1e6, rec.status, r.Method, r.URL.Path)
		return
	}
}

func (mdw *Middlewares) ErrorHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				// if error occured
				log.Println("Panic caught:", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next(w, r)
		return
	}
}
