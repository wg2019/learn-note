// Package factory 测试
package factory_test

import (
	"github.com/wg2019/design-pattern/creation/factory"
	"testing"
)

// TestFactoryPattern 测试工厂模式.
func TestFactoryPattern(t *testing.T) {
	new(factory.AFactory).Manufacture().Show()
	new(factory.BFactory).Manufacture().Show()
}
