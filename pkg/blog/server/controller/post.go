package controller

import (
	"net/http"
)

type PostController struct{}

func (ctrl *PostController) GetPostList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey"))
}

// func (ctrl *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	util.ChkErr(err)
//
// 	var post *model.Post
// 	err = json.Unmarshal(body, post)
// 	util.ChkErr(err)
//
// }
