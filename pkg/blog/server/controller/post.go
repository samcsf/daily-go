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
		ctrl.GetPosts(w, r)
	case http.MethodPost:
		ctrl.CreatePost(w, r)
	case http.MethodPut:
		ctrl.UpdatePost(w, r)
	case http.MethodDelete:
		ctrl.DelPost(w, r)
	}
}

func (ctrl *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := service.Post.GetPosts()
	util.ChkErr(err)

	bytes, err := json.Marshal(posts)
	util.ChkErr(err)

	w.Write(bytes)
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

func (ctrl *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	util.ChkErr(err)

	var post *model.Post
	err = json.Unmarshal(body, &post)
	util.ChkErr(err)

	if post.Id == "" {
		http.Error(w, "Post id is required", http.StatusBadRequest)
		return
	}

	service.Post.UpdatePost(post)
	bytes, err := json.Marshal(map[string]int{"ok": 1})
	util.ChkErr(err)

	w.Write(bytes)
}

func (ctrl *PostController) DelPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	util.ChkErr(err)

	var post *model.Post
	err = json.Unmarshal(body, &post)
	util.ChkErr(err)

	if post.Id == "" {
		http.Error(w, "Post id is required", http.StatusBadRequest)
		return
	}

	service.Post.DelPost(post)
	bytes, err := json.Marshal(map[string]int{"ok": 1})
	util.ChkErr(err)

	w.Write(bytes)
}
