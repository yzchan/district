package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
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
		//fmt.Println(i+1, section, province, city, areacode, postcode, isp, simcard)
		if code != "" {
			fmt.Println(code, strings.Replace(name, string([]byte{0xC2, 0xA0}), "*", -1))
		}

	})
	//for i, r := range records {
	//	fmt.Println(i, r)
	//}

}
