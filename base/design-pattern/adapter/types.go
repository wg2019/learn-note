// Package adapter 定义
package adapter

// Restaurant 饭店 目标对象
type Restaurant interface {
	Food()
}

// Shop 超市 兼容对象
type Shop interface {
	Bread()
}
