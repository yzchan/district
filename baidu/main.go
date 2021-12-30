package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/xuri/excelize/v2"
)

var (
	url = "https://mapopen-pub-webserviceapi.bj.bcebos.com/geocoding/Township_Area_A_202104.xlsx"
)

func main() {
	// 下载xlsx数据
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("Township_Area_A_202104.xlsx")
	if err != nil {
		panic(err)
	}
	io.Copy(file, res.Body)
	file.Close()

	// 读取xlsx数据
	f, err := excelize.OpenFile("Township_Area_A_202104.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	//defer func() {
	//	// Close the spreadsheet.
	//	if err := f.Close(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	// Get value from cell by given worksheet name and axis.
	//cell, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
