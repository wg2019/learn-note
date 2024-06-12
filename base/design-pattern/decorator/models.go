// Package decorator 实现
package decorator

import "fmt"

// Liberation 解放
type Liberation struct {
}

// NewLiberation 解放
func NewLiberation() BigCar {
	return new(Liberation)
}

// Checklist 清单
func (l *Liberation) Checklist() []string {
	checkList := make([]string, 0, 3)
	checkList = append(checkList, "我是解放大卡车!")
	return checkList
}

// FAW 一汽
type FAW struct {
	bigCar BigCar
	times  int
}

// Checklist 一汽改进清单
func (f *FAW) Checklist() []string {
	return append(f.bigCar.Checklist(), fmt.Sprintf("第%d次改进", f.times))
}

// NewFAW 获取装饰器
func NewFAW(car BigCar, times int) BigCar {
	f := new(FAW)
	f.bigCar = car
	f.times = times
	return f
}
