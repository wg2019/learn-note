// Package builder 测试
package builder

import "testing"

// TestMingDirector 测试小明
func TestMingDirector(t *testing.T) {
	director := MingDirector()
	count := director.GetCount()
	productList := director.GetProductList()
	t.Logf("count: %d\n", count)
	for _, product := range productList {
		t.Logf("product Name: %s, Color: %s, Size: %d\n",
			product.Name(), product.Color(), product.Size())
	}
}

// TestKingDirector 测试小王
func TestKingDirector(t *testing.T) {
	director := KingDirector()
	count := director.GetCount()
	productList := director.GetProductList()
	t.Logf("count: %d\n", count)
	for _, product := range productList {
		t.Logf("product Name: %s, Color: %s, Size: %d\n",
			product.Name(), product.Color(), product.Size())
	}
}
