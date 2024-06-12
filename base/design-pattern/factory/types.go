// Package factory 定义
package factory

// Factory 工厂.
type Factory interface {
	Manufacture() Product
}

// Product 产品.
type Product interface {
	Show()
}
