// Package factory 实现
package factory

import "fmt"

// AFactory 工厂A.
type AFactory struct {
}

// AProduct 产品A.
type AProduct struct {
}

// Show 内容.
func (a *AProduct) Show() {
	fmt.Printf("name: %s\n", "AProduct")
}

// Manufacture 获取产品.
func (a *AFactory) Manufacture() Product {
	fmt.Printf("name: %s\n", "AFactory")
	return new(AProduct)
}

// BFactory 工厂B.
type BFactory struct {
}

// BProduct 产品B.
type BProduct struct {
}

// Show 内容.
func (a *BProduct) Show() {
	fmt.Printf("name: %s\n", "BProduct")
}

// Manufacture 获取产品.
func (a *BFactory) Manufacture() Product {
	fmt.Printf("name: %s\n", "BFactory")
	return new(BProduct)
}
