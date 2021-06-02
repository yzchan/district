package district

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Base struct {
	AdCode   int
	Name     string
	AreaCode string
	PostCode string
}

type Province struct {
	Base
	ShortName string
	Cities    []City
}

type City struct {
	Base
	Areas []Area
}
type Area struct {
	Base
}

var Maps map[int]string

var Instance *district

type district struct {
	provinces  map[int]string
	cities     map[int]string
	areas      map[int]string
	provCities map[int][]int
	cityAreas  map[int][]int
}

func init() {
	f, _ := os.Open("data/3-level.csv")
	r := csv.NewReader(f)
	data, _ := r.ReadAll()
	Maps = make(map[int]string, 0)
	for _, line := range data[1:] {
		code, _ := strconv.Atoi(line[0])
		Maps[code] = line[1]
	}

	Instance = &district{}
	Instance.provinces = make(map[int]string, 0)
	Instance.cities = make(map[int]string, 0)
	Instance.areas = make(map[int]string, 0)
	Instance.provCities = make(map[int][]int, 0)
	Instance.cityAreas = make(map[int][]int, 0)
	formatData()
}

func formatData() {
	for code, name := range Maps {
		if code%10000 == 0 {
			Instance.provinces[code] = name
			Instance.provCities[code] = make([]int, 0)
		} else if code%100 == 0 && code%10000 > 0 {
			Instance.cities[code] = name
			Instance.cityAreas[code] = make([]int, 0)
		} else if code%100 > 0 {
			Instance.areas[code] = name
		}
	}
	for code := range Instance.cities {
		provCode := getProvinceCode(code)
		Instance.provCities[provCode] = append(Instance.provCities[provCode], code)
	}
	for code := range Instance.areas {
		cityCode := getCityCode(code)
		Instance.cityAreas[cityCode] = append(Instance.cityAreas[cityCode], code)
	}
}

func getProvinceCode(code int) int {
	return code / 10000 * 10000
}

func getCityCode(code int) int {
	return code / 100 * 100
}

// Search 根据关键词搜索地区
func (d *district) Search(keyword string) (result map[int]string) {
	result = make(map[int]string, 0)
	for code, name := range Maps {
		if strings.Contains(name, keyword) {
			result[code] = name
		}
	}
	return
}

// ShortNames 获取省级行政区的简称
func (d *district) ShortNames(code int) (primary string, secondary string) {
	province, exists := d.provinces[code]
	if exists {
		names := ProvShortName[province]
		primary = names[0]
		if len(names) > 1 {
			secondary = names[1]
		}
	}
	return
}

// IsDirectCity 是否直辖市
func (d *district) IsDirectCity(code int) bool {
	return code == 110000 || code == 120000 || code == 310000 || code == 500000
}

// IsProvince 是否省级行政单位
func (d *district) IsProvince(code int) bool {
	_, ok := d.provinces[code]
	return ok
}

// Provinces 返回所有管辖城市
func (d *district) Provinces() map[int]string {
	provinces := make(map[int]string, 0)
	for code, name := range d.provinces {
		provinces[code] = name
	}
	return provinces
}

// Cities 返回指定省份的下属城市（直辖市会直接返回其下属区）
func (d *district) Cities(provCode int) map[int]string {
	cities := make(map[int]string, 0)
	for _, code := range d.provCities[provCode] {
		cities[code] = d.cities[code]
	}
	return cities
}

// Areas 返回指定城市的所有区（支持传入直辖市的行政区划代码）
func (d *district) Areas(cityCode int) map[int]string {
	//if d.IsDirectCity(cityCode) { // 兼容对直辖市的处理
	//	return d.Cities(cityCode)
	//}
	areas := make(map[int]string, 0)
	for _, code := range d.cityAreas[cityCode] {
		areas[code] = d.areas[code]
	}
	return areas
}
