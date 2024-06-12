// Package prototype 实现
package prototype

// Clone 克隆方法
func (p *product) Clone() *product {
	models := *p
	return &models
}

// product 原型产品
type product struct {
	Name   string
	Color  string
	Size   int
	Weight int
}

// 获得产品
func GetProduct() *product {
	p := new(product)
	return p
}
