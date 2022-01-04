package main

import (
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strings"
)

var url = "http://www.mca.gov.cn/article/sj/xzqh/2020/20201201.html"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	f, err := os.Create("data.csv")
	w := csv.NewWriter(f)
	_ = w.Write([]string{"Code", "Name"})
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	//var records []PhoneSection
	// Find the review items
	dom.Find("table tr").Each(func(i int, s *goquery.Selection) {
		code := s.Find("td").Eq(1).Text()
		//pk, _ := strconv.Atoi(section)
		name := s.Find("td").Eq(2).Text()
		if code != "" {
			//strings.Replace(name, string([]byte{0xC2, 0xA0}), "", -1)
			name = strings.TrimSpace(name)
			_ = w.Write([]string{code, name})
		}
	})
	w.Flush()

}
