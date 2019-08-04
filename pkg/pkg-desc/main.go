// pkg like github.com/samcsf/pkg-refs
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {
	idvFlag := flag.Bool("f", false, "Fetch individual repo")
	flag.Parse()

	if len(os.Args) < 2 || (*idvFlag && len(os.Args) < 3) {
		log.Fatalln("Pkg name must provided")
	}

	if *idvFlag {
		pkg := os.Args[2]
		parseURL(pkg)
		return
	}

	pkg := os.Args[1]
	parseFile(pkg)
}

func parseURL(pkg string) {
	fetchRepo(pkg)
}

func fetchRepo(pkg string) {
	match, err := regexp.MatchString("^https?://", pkg)
	chkErr(err)

	if !match {
		pkg = "https://" + pkg
	}

	re, err := regexp.Compile(`[^/][a-zA-Z-_]*$`)
	chkErr(err)

	pkgName := re.FindString(pkg)
	resp, err := http.Get(pkg)
	chkErr(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	chkErr(err)

	doc.Find("span.text-gray-dark.mr-2").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		text = strings.TrimSpace(text)
		color.Cyan("%s", pkgName)
		fmt.Printf("  %s\n", text)
	})
}

func parseFile(pkg string) {
	info, err := os.Stat(pkg)
	if os.IsNotExist(err) {
		log.Fatalln("Path not exists")
	}

	if info.IsDir() {
		pkg = path.Join(pkg, "go.mod")
	}

	file, err := os.OpenFile(pkg, os.O_RDONLY, 0666)
	chkErr(err)

	fileContent, err := ioutil.ReadAll(file)
	chkErr(err)
	fileString := string(fileContent)

	// parse the repos
	githubPtn, _ := regexp.Compile(`github.com\S*`)
	repos := githubPtn.FindAllString(fileString, -1)

	var wg sync.WaitGroup
	wg.Add(len(repos))
	for _, repo := range repos {
		go func(repo string) {
			fetchRepo(repo)
			wg.Done()
		}(repo)
	}

	wg.Wait()
}

func chkErr(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}
