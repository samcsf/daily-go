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
	d := cps.Compose(ctrl.Mdw.Duration)
	http.HandleFunc("/", d(ctrl.Post.GetPostList))
	http.HandleFunc("/post", d(ctrl.Post.CreatePost))

	log.Println("Blog server started on port", config.ServerPort)
	err := http.ListenAndServe(config.ServerPort, nil)
	util.ChkErr(err)
}
