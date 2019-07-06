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

func (ps *PostService) SavePost(post *model.Post) (sql.Result, error) {
	stmt, err := model.MySqlDB.Prepare("INSERT INTO posts VALUES (?,?,?,?,?)")
	util.ChkErr(err)
	now := time.Now()
	res, err := stmt.Exec(nil, post.Title, post.Content, now, now)
	util.ChkErr(err)
	return res, nil
}
