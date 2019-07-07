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
