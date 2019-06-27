package crawler

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
)

var url = "https://www.douyu.com/"

func chkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// 斗鱼热门视频分类
func FetchDouyuHotCategory() {
	resp, err := http.Get(url)
	chkErr(err)
	defer resp.Body.Close()

	// load body, parse dom with lib goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	chkErr(err)
	var jsData string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if i == 5 {
			jsData = text
		}
	})

	// load sript and parse with lib otto (arrow function not supported!)
	vm := otto.New()
	vm.Run(jsData)
	val, err := vm.Run("$mainData[6].ds.slice(0, 6).sort(function(a, b){ return b.cnt - a.cnt }).map(function(i){ return i.cn}).join('|')")
	chkErr(err)
	log.Printf("斗鱼热门分类: %s", val)
}
