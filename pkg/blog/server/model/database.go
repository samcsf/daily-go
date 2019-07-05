package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samcsf/daily-go/pkg/blog/server/config"
	"github.com/samcsf/daily-go/pkg/util"
)

var MySqlDB *sql.DB

func init() {
	db, err := sql.Open("mysql", config.MySqlUrlStr)
	util.ChkErr(err)
	MySqlDB = db
}
