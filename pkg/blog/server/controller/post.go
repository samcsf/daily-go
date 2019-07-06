package controller

import (
	"encoding/json"
	"github.com/samcsf/daily-go/pkg/blog/server/model"
	"github.com/samcsf/daily-go/pkg/blog/server/service"
	"github.com/samcsf/daily-go/pkg/util"
	"io/ioutil"
	// "log"
	"net/http"
)

type PostController struct{}

func (ctrl *PostController) GetPostList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey"))
}

func (ctrl *PostController) HandlePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
		ctrl.CreatePost(w, r)
	case http.MethodPut:
	case http.MethodDelete:
	}
}

func (ctrl *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	util.ChkErr(err)

	var post *model.Post
	err = json.Unmarshal(body, &post)
	util.ChkErr(err)

	service.Post.SavePost(post)
	bytes, err := json.Marshal(map[string]int{"ok": 1})
	util.ChkErr(err)

	w.Write(bytes)
}
