// Package builder 定义
package builder

// Product 产品.
type Product interface {
	Name() string
	Color() string
	Size() int
}

// Factory 工厂.
type Factory interface {
	GetCount() int
	GetProductList() []Product
	AddProduct(Product)
}
