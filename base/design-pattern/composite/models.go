package composite

import "fmt"

// province 省份
type province struct {
	name     string
	cityList []Composite
}

// AddComposite 添加组织
func (h *province) AddComposite(city Composite) {
	h.cityList = append(h.cityList, city)
}

// Features 功能
func (h *province) Features() {
	fmt.Printf("省份名称：%s\n", h.name)
	for _, city := range h.cityList {
		fmt.Printf("\t")
		city.Features()
	}
}

// city 城市
type city struct {
	name       string
	regionList []Composite
}

// AddComposite 添加组织
func (c *city) AddComposite(region Composite) {
	c.regionList = append(c.regionList, region)
}

// Features 功能
func (c *city) Features() {
	fmt.Printf("城市名称：%s\n", c.name)
	for _, region := range c.regionList {
		fmt.Printf("\t\t")
		region.Features()
	}
}

// region 地区
type region struct {
	name string
}

// AddComposite 添加组织
func (r *region) AddComposite(Composite) {
}

// Features 功能
func (r *region) Features() {
	fmt.Printf("地区名称：%s\n", r.name)
}

// NewProvince 获取省份
func NewProvince(name string) Composite {
	p := new(province)
	p.name = name
	p.cityList = make([]Composite, 0, 3)
	return p
}

// NewCity 获取城市
func NewCity(name string) Composite {
	c := new(city)
	c.name = name
	c.regionList = make([]Composite, 0, 3)
	return c
}

// NewRegion 获取地区
func NewRegion(name string) Composite {
	r := new(region)
	r.name = name
	return r
}
