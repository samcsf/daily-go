package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samcsf/daily-go/pkg/blog/server/model"
	"github.com/samcsf/daily-go/pkg/util"
	"time"
)

type PostService struct{}

/*
   Id
   Title
   Content
   Create_at
   Modified_at
*/

func (ps *PostService) GetPosts() ([]*model.Post, error) {
	res, err := model.MySqlDB.Query("SELECT * FROM posts")
	util.ChkErr(err)

	var posts []*model.Post
	for res.Next() {
		post := &model.Post{}
		var cDate, mDate string
		err = res.Scan(&post.Id, &post.Title, &post.Content, &cDate, &mDate)
		util.ChkErr(err)

		loc, _ := time.LoadLocation("Asia/Shanghai")
		post.Create_at, err = time.ParseInLocation("2006-01-02 15:04:05", cDate, loc)
		util.ChkErr(err)
		post.Modified_at, err = time.ParseInLocation("2006-01-02 15:04:05", mDate, loc)
		util.ChkErr(err)

		posts = append(posts, post)
	}

	return posts, nil
}

func (ps *PostService) UpdatePost(post *model.Post) (sql.Result, error) {
	stmt, err := model.MySqlDB.Prepare("UPDATE posts SET title=?, content=?, modified_at=? WHERE id=?")
	util.ChkErr(err)
	now := time.Now()
	res, err := stmt.Exec(post.Title, post.Content, now, post.Id)
	util.ChkErr(err)
	return res, nil
}

func (ps *PostService) SavePost(post *model.Post) (sql.Result, error) {
	stmt, err := model.MySqlDB.Prepare("INSERT INTO posts VALUES (?,?,?,?,?)")
	util.ChkErr(err)
	now := time.Now()
	res, err := stmt.Exec(nil, post.Title, post.Content, now, now)
	util.ChkErr(err)
	return res, nil
}

func (ps *PostService) DelPost(post *model.Post) (sql.Result, error) {
	stmt, err := model.MySqlDB.Prepare("DELETE FROM posts WHERE id=?")
	util.ChkErr(err)
	res, err := stmt.Exec(post.Id)
	util.ChkErr(err)
	return res, nil
}
