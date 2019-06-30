package main

import (
	"expvar"
	"fmt"
	"net/http"
)

// NewXXX will register new var, internally invoke Publish()
var visited = expvar.NewInt("visited")

type MySystemExpString string

var iamNotExposedDefault = "invisble"

func (str MySystemExpString) String() string {
	return iamNotExposedDefault
}

func handler(w http.ResponseWriter, r *http.Request) {
	visited.Add(1)
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Now the visit count is %d", visited.Value())
}

// invoke Do with iterate each kv and process (safe)
func listHandler(w http.ResponseWriter, r *http.Request) {
	var strs []string
	expvar.Do(func(kv expvar.KeyValue) {
		strs = append(strs, kv.Key+" "+kv.Value.String())
	})
	res := ""
	for _, str := range strs {
		res += (str + "\n")
	}
	fmt.Fprintf(w, "%s", res)
}

func exposeHandler(w http.ResponseWriter, r *http.Request) {
	// Publish any var to debug vars will alias(or you can use same one)
	expvar.Publish("nowIsExposed", MySystemExpString(""))
	fmt.Fprintf(w, "%s", "ok")
}

func visibleHandler(w http.ResponseWriter, r *http.Request) {
	iamNotExposedDefault = "visible"
	fmt.Fprintf(w, "%s", "ok")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/expose", exposeHandler)
	http.HandleFunc("/visible", visibleHandler)
	// add alias path
	http.Handle("/metrics", expvar.Handler())
	http.ListenAndServe(":8888", nil)
}
