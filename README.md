中国行政区划数据（2020年11月）
-----

我国的行政区划共分为4个等级：省级行政区 (Province)、地级行政区(City)、县级行政区 (Country)、乡级行政区(Town)。（截止2021年5月的数据）

[省级行政区](https://baike.baidu.com/item/省级行政区/4805340) 为一级行政区，包括省、自治区、直辖市、特别行政区。 

[地级行政区](https://baike.baidu.com/item/地级行政区/5628580) 为二级行政区，包括地级市、地区、自治州、盟。

[县级行政区](https://baike.baidu.com/item/县级行政区/1660163) 为三级行政区，包括市辖区、县级市、县、自治县、旗、自治旗、特区、林区。

[乡级行政区](https://baike.baidu.com/item/乡级行政区/8466613) 为四级行政区，包括街道、镇、乡、民族乡、苏木、民族苏木、县辖区。

本项目数据基于以下来源制作：

- 官方数据： [2020年11月中华人民共和国县以上行政区划代码](http://www.mca.gov.cn/article/sj/xzqh/2020/)
- 百度地图： [百度地图行政区划adcode映射表【更新至21年04月】](https://lbsyun.baidu.com/index.php?title=open/dev-res)

其中官方行政区划数据只包含了一二三级行政区数据，百度地图提供的数据补充了第四级行政区的数据以及港澳台的二三级行政区数据（目前不含港澳台地区四级行政区数据）。

大部分地址会包含全部四级行政区，比如"江苏省常州市新北区三井街道"。但是直辖市和直辖县级市例外。

|  省级   | 地级     | 县级   |乡级  |
|  ----  | ----     | ----  | ----  |
| 上海    | -       | 浦东新区  | 金桥镇 |

上表为"上海浦东新区金桥镇"的行政区划。该地址不含低级行政区。

|  省级   | 地级     | 县级   |乡级  |
|  ----  | ----     | ----  | ----  |
| 新疆    | 乌鲁木齐  | 天山区    | 红雁街道 |
| 新疆    | -       | 石河子市   | 新城街道 |

上表中石河子市为新疆维吾尔自治区直辖县级市，直接隶属新疆维吾尔自治区。所以该地址也不含低级行政区，与直辖市表现的一样。

## About

本项目包含三级行政区数据和四级行政区数据两部分。三级行政区数据基于官方数据制作，使用map存储区划代码和区划名称，查询效率高。

四级行政区划数据基于官方数据和百度地图制作而成，采用数据库存储。

## Installation

```shell
go get -u github.com/yzchan/cn-district
```

## Quickstart

```golang
package main

import (
	"fmt"
	"github.com/yzchan/cn-district/level3"
)

func main() {
	// 直接调用原始数据
	fmt.Println(level3.DistrictMap)
	fmt.Println(level3.ProvShortName)

	// 创建单例实例
	var d = level3.Instance

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

- 增加邮编和区号数据
- 增加更多数据格式化方法