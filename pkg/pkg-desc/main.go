package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Pkg name must provided")
	}
	pkg := "https://" + os.Args[1]

	resp, err := http.Get(pkg)
	chkErr(err)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	chkErr(err)

	doc.Find("span.text-gray-dark.mr-2").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		log.Println(text)
	})
}

func chkErr(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}
