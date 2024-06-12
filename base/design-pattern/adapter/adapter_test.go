// Package adapter 测试
package adapter

import "testing"

// TestAdapter_Food 测试
func TestAdapter_Food(t *testing.T) {
	a := new(Adapter)
	a.WumartSupermarket.LongBread = "长面包"
	a.Food()
}
