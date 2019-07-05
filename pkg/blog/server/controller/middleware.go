package controller

import (
	"log"
	"net/http"
	"time"
)

type Middlewares struct{}

func (mdw *Middlewares) Duration(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			log.Printf("| %dms | %s %s", end.Sub(start).Nanoseconds()/1e6, r.Method, r.URL.Path)
		}()
		next(w, r)
		return
	}
}
