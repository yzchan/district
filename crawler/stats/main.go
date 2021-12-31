package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"golang.org/x/text/encoding/simplifiedchinese"
	"time"
)

func main() {
	enc := simplifiedchinese.GBK.NewDecoder()
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.stats.gov.cn"),
	)
	_ = c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1, Delay: time.Millisecond * 1000})

	c.OnHTML(".villagetr", func(e *colly.HTMLElement) {
		code := e.DOM.Find("td").Eq(0).Text()
		cate := e.DOM.Find("td").Eq(1).Text()
		name := e.DOM.Find("td").Eq(2).Text()
		decodedName, _ := enc.String(name)
		fmt.Println(code, cate, decodedName)
	})

	c.OnHTML(".countytr,.towntr,.citytr", func(e *colly.HTMLElement) {
		code := e.DOM.Find("td").Eq(0).Text()
		name := e.DOM.Find("td").Eq(1).Text()
		decodedName, _ := enc.String(name)
		fmt.Println(e.Attr("class"), code, decodedName)
	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/index.html") // 入口地址
}
