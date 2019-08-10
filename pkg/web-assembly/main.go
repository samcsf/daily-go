package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/", wasmHanlerMid(http.FileServer(http.Dir("./pkg/web-assembly/html/"))))
	log.Println("Server start listening 8080")
	http.ListenAndServe(":8080", nil)
}

// Go在web-assembly中其实就是取代js的位置，服务端所需要做的就是设置正确的content-type让浏览器识别
func wasmHanlerMid(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-type", "application/wasm")
		}
		h.ServeHTTP(w, r)
	})
}
