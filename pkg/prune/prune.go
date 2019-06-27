/*
	受TJ大神node-prune启发写个简单的go-prune，清除vendor中不必要的文件
*/
package prune

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
)

var PruneList = []string{
	`^\..*`,
	`.*\.md`,
	`.*\.markdown`,
	`makefile`,
	`dockerfile`,
	`AUTHORS`,
	`LICENSE`,
	`CONTRIBUTORS`,
	`PATENTS`,
}

var totalCleanSize int64
var totalCleanCount int

func chkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func Prune(root string) {
	// find root/vendor folder
	target := path.Join(root, "vendor")
	// walk throught the folder
	// regexp test and remove
	walk(target, handle)
	log.Println("Total", totalCleanCount, "files cleaned")
	log.Println("Saved", totalCleanSize, "bytes")
}

func handle(baseDir string, info os.FileInfo) {
	fullPath := path.Join(baseDir, info.Name())
	// log.Println(fullPath, info.Size(), "bytes")
	for _, pattern := range PruneList {
		isMatch, err := regexp.MatchString("(?i)"+pattern, info.Name())
		chkErr(err)
		if isMatch {
			log.Println("X", fullPath)
			err = os.Remove(fullPath)
			chkErr(err)
			totalCleanCount++
			totalCleanSize += info.Size()
			break
		}
	}
}

func walk(filePath string, fn func(string, os.FileInfo)) {
	log.Println("Start to scan from ", filePath)
	files, err := ioutil.ReadDir(filePath)
	chkErr(err)
	for _, file := range files {
		fullPath := path.Join(filePath, file.Name())
		if file.IsDir() {
			walk(fullPath, fn)
		} else {
			fn(filePath, file)
		}
	}
}
