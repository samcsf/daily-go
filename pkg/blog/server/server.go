package main

import (
	"github.com/samcsf/daily-go/pkg/blog/server/config"
	ctrl "github.com/samcsf/daily-go/pkg/blog/server/controller"
	cps "github.com/samcsf/daily-go/pkg/compose"
	"github.com/samcsf/daily-go/pkg/util"
	"log"
	"net/http"
)

func main() {
	d := cps.Compose(ctrl.Mdw.ErrorHandle, ctrl.Mdw.Duration)

	// static server
	fileHandler := http.FileServer(http.Dir("./pkg/blog/server/public"))

	http.Handle("/", fileHandler)
	http.HandleFunc("/post", d(ctrl.Post.HandlePost))

	config.ReadFromJson()

	log.Println("Blog server started on port", config.ServerPort)
	err := http.ListenAndServe(config.ServerPort, nil)
	util.ChkErr(err)
}
