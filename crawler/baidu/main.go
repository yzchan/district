package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"net/http"
)

var (
	url = "https://mapopen-pub-webserviceapi.bj.bcebos.com/geocoding/Township_Area_A_202104.xlsx"
)

func main() {
	// 下载xlsx数据
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// 读取xlsx数据
	f, err := excelize.OpenReader(resp.Body)
	if err != nil {
		panic(err)
	}

	// 读取Sheet1所有行
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
