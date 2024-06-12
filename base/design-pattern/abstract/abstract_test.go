package abstract

import (
	"fmt"
	"testing"
)

// TestFactoryPattern 测试工厂模式.
func TestAbstractPattern(t *testing.T) {
	shanDong := new(ShanDong)
	fmt.Printf("product name: %s\n", shanDong.GetBanana().Name())
	fmt.Printf("product name: %s\n", shanDong.GetApple().Name())
	yunNan := new(YunNan)
	fmt.Printf("product name: %s\n", yunNan.GetBanana().Name())
	fmt.Printf("product name: %s\n", yunNan.GetApple().Name())
}
