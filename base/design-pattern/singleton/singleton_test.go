// Package singleton 测试
package singleton

import (
	"testing"
)

// TestGetLazyInstance 懒汉式
func TestGetLazyInstance(t *testing.T) {
	instance := GetLazyInstance()
	t.Logf("name: %s", instance.name)
}

// TestGetLazyInstance 懒汉式
func TestGetHungryInstance(t *testing.T) {
	instance := GetHungryInstance()
	t.Logf("name: %s", instance.name)
}

// TestGetDoubleInstance 双重检查机制
func TestGetDoubleInstance(t *testing.T) {
	instance := GetDoubleInstance()
	t.Logf("name: %s", instance.name)
}

// TestGetOnceInstance sync.Once
func TestGetOnceInstance(t *testing.T) {
	instance := GetOnceInstance()
	t.Logf("name: %s", instance.name)
}
