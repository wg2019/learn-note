package template

import "fmt"

// ChineseFood 中餐
type ChineseFood struct{}

// Name 名称
func (c *ChineseFood) Name() {
	fmt.Printf("---中餐---\n")
}

// StapleFood 主食
func (c *ChineseFood) StapleFood() {
	fmt.Printf("主食：米饭\n")
}

// Dish 菜品
func (c *ChineseFood) Dish() {
	fmt.Printf("菜品：土豆丝\n")
}

// Drink 饮料
func (c *ChineseFood) Drink() {
	fmt.Printf("饮料：绿豆粥\n")
}

// WesternFood 西餐
type WesternFood struct{}

// Name 名称
func (w *WesternFood) Name() {
	fmt.Printf("---西餐---\n")
}

// StapleFood 主食
func (w *WesternFood) StapleFood() {
	fmt.Printf("主食：面包\n")
}

// Dish 菜品
func (w *WesternFood) Dish() {
	fmt.Printf("菜品：牛排\n")
}

// Drink 饮料
func (w *WesternFood) Drink() {
	fmt.Printf("饮料：牛奶\n")
}
