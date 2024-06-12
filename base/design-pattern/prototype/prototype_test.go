package prototype

import (
	"fmt"
	"testing"
)

// TestGetProduct 测试
func TestGetProduct(t *testing.T) {
	productA := GetProduct()
	productA.Name = "产品A"
	productA.Color = "黄色"
	productA.Size = 10
	productA.Weight = 200
	productB := productA.Clone()
	productB.Name = "产品B"
	productB.Color = "红色"
	productC := productB.Clone()
	productC.Name = "产品C"
	productC.Size = 20
	fmt.Printf("名称: %s, 颜色: %s, 型号: %d, 重量: %d\n",
		productA.Name, productA.Color, productA.Size, productA.Weight)
	fmt.Printf("名称: %s, 颜色: %s, 型号: %d, 重量: %d\n",
		productB.Name, productB.Color, productB.Size, productB.Weight)
	fmt.Printf("名称: %s, 颜色: %s, 型号: %d, 重量: %d\n",
		productC.Name, productC.Color, productC.Size, productC.Weight)
}
