中国行政区划数据（2020年11月）
-----

## About

本项目数据来源：[2020年11月中华人民共和国县以上行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/2020/)

地址分三级：省级行政区划（一级）、市级行政区划（二级）、区级行政区划（三级）

一级行政区划包括4中类型：省、直辖市、自治区、特别行政区。

## Installation

```shell
go get -u github.com/yzchan/cn-district
```

## Quickstart

```golang
package main

import (
	"fmt"
	district "github.com/yzchan/cn-district"
	"github.com/yzchan/cn-district/data"
)

func main() {
	// 直接调用原始数据
	fmt.Println(data.DistrictMap)
	fmt.Println(data.ProvShortName)

	// 创建单例实例
	var d = district.Instance

	// 获取所有省份数据
	provinces := d.Provinces()
	fmt.Println(provinces)

	// 获取省级行政区下辖市（直辖市会返回下辖区）
	cities := d.Cities(320000)
	fmt.Println(cities)

	// 获取市级行政区划下辖区
	// 直辖市这边做了特别处理，也接受直辖市区划代码。d.Cities(310000)以及d.Areas(310000)返回数据跟完全相同
	areas := d.Areas(320400)
	fmt.Println(areas)

	// 判断是否直辖市
	fmt.Println(d.IsDirectCity(310000))

	// 判断是否省级行政区划代码
	fmt.Println(d.IsProvince(310000))

	// 根据关键词搜索地区
	result := d.Search("上海")
	fmt.Println(result)

	// 获取省份简称 例如四川[510000]返回 "川" "蜀"
	primary, secondary := d.ShortNames(510000)
	fmt.Println(primary, secondary)
}

```

## Features