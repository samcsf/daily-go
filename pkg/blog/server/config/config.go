package config

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"os"
)

var ServerPort = ":8888"

var LogTo = "stdout"
var MySqlUser = "root"
var MySqlPass = ""
var MySqlAddress = ""
var MySqlDB = "goblog"
var MySqlUrlStr = fmt.Sprintf("%s:%s@%s/%s", MySqlUser, MySqlPass, MySqlAddress, MySqlDB)

func init() {
	readFromJson()
	port := os.Getenv("PORT")
	if port != "" {
		ServerPort = ":" + port
	}
	mapp := config.Map()["service"].(map[string]interface{})
	if val, ok := mapp["port"]; ok {
		ServerPort = val
	}
	logto := os.Getenv("LOGTO")
	if logto != "" {
		LogTo = logto
	}
}

func readFromJson() {
	config.LoadFile("./pkg/blog/server/config.json")
	config.Get("service")
	fmt.Println(config.Map()["service"])
}
