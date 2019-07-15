package controller

import (
	"encoding/json"
	"github.com/samcsf/daily-go/pkg/blog/server/model"
	"github.com/samcsf/daily-go/pkg/blog/server/service"
	"github.com/samcsf/daily-go/pkg/util"
	"github.com/thedevsaddam/govalidator"
	"io/ioutil"
	"net/http"
)

type PostController struct{}

func (ctrl *PostController) Ping(w http.ResponseWriter, r *http.Request) {
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

	err = json.NewEncoder(w).Encode(posts)
	util.ChkErr(err)
}

func (ctrl *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"title":   []string{"required"},
		"content": []string{"required"},
	}
	post := &model.Post{}
	opts := govalidator.Options{
		Request: r,
		Rules:   rules,
		Data:    post,
	}
	e := govalidator.New(opts).ValidateJSON()
	if len(e) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		erx := util.H{"validationError": e}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(erx)
		return
	}

	service.Post.SavePost(post)
	err := json.NewEncoder(w).Encode(util.H{"ok": 1})
	util.ChkErr(err)
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
	err = json.NewEncoder(w).Encode(util.H{"ok": 1})
	util.ChkErr(err)
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
	err = json.NewEncoder(w).Encode(util.H{"ok": 1})
	util.ChkErr(err)
}
