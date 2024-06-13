// Package factory 测试
package factory

import (
	"testing"
)

// TestFactoryPattern 测试工厂模式.
func TestFactoryPattern(t *testing.T) {
	GetFactory("AFactory").Manufacture().Show()
	GetFactory("BFactory").Manufacture().Show()
}
