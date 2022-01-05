package main

import (
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var url = "http://www.mca.gov.cn/article/sj/xzqh/2020/20201201.html"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	f, err := os.Create("mca.csv")
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
		_, ok := strconv.Atoi(code)
		if ok != nil {
			return
		}
		//strings.Replace(name, string([]byte{0xC2, 0xA0}), "", -1)
		name := strings.TrimSpace(s.Find("td").Eq(2).Text())
		w.Write([]string{code, name})

	})
	w.Flush()

}
