package config

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"os"
)

var ServerPort = ":8888"

var MySqlUser = "root"
var MySqlPass = ""
var MySqlAddress = ""
var MySqlDB = "goblog"
var MySqlUrlStr = fmt.Sprintf("%s:%s@%s/%s", MySqlUser, MySqlPass, MySqlAddress, MySqlDB)

func init() {
	port := os.Getenv("PORT")
	if port != "" {
		ServerPort = ":" + port
	}
}

func ReadFromJson() {
	config.LoadFile("../config.json")
	fmt.Println(config.Get("service", "port"))
}
