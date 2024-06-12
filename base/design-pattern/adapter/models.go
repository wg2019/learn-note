// Package adapter 实现
package adapter

import "fmt"

// PeaceRestaurant 和平饭店
type PeaceRestaurant struct {
	ChineseFood string
}

// Food 食物
func (p *PeaceRestaurant) Food() {
	fmt.Printf("和平饭店：%s\n", p.ChineseFood)
}

// WumartSupermarket 物美超市
type WumartSupermarket struct {
	LongBread string
}

// Bread 面包
func (w *WumartSupermarket) Bread() {
	fmt.Printf("物美超市：%s\n", w.LongBread)
}

// Adapter 适配器
type Adapter struct {
	WumartSupermarket
}

// Food 食物
func (a *Adapter) Food() {
	a.WumartSupermarket.Bread()
}
