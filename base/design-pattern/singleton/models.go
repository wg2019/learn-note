// Package singleton 实现
package singleton

import "sync"

// lazyInstance 实例
var lazyInstance *singleton

// GetLazyInstance 懒汉式
func GetLazyInstance() *singleton {
	if lazyInstance == nil {
		lazyInstance = new(singleton)
		lazyInstance.name = "懒汉式"
	}
	return lazyInstance
}

// hungryInstance 实例
var hungryInstance *singleton

func init() {
	hungryInstance = new(singleton)
	hungryInstance.name = "饿汉式"
}

// GetHungryInstance 饿汉式
func GetHungryInstance() *singleton {
	return hungryInstance
}

// doubleInstance 实例
var doubleInstance *singleton

// 加锁
var mux sync.Mutex

// GetDoubleInstance 双重检查机制
func GetDoubleInstance() *singleton {
	if doubleInstance == nil {
		mux.Lock()
		defer mux.Unlock()
		if doubleInstance == nil {
			doubleInstance = new(singleton)
			doubleInstance.name = "双重检查机制"
		}
	}
	return doubleInstance
}

// 实例
var onceInstance *singleton
var once sync.Once

// GetDoubleInstance sync.Once
func GetOnceInstance() *singleton {
	once.Do(func() {
		onceInstance = new(singleton)
		onceInstance.name = "sync.Once"
	})
	return onceInstance
}
