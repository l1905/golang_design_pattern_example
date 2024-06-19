package singleton

import (
	"fmt"
	"sync"
)

// 创建型模式-单例模式
// Singleton 是单例模式的结构
type Singleton struct {
	// 可以定义一些字段
}

// 单例对象
var instance *Singleton

// sync.Once 用于确保单例对象只被创建一次
var once sync.Once

// GetInstance 返回单例对象
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		fmt.Println("创建单例对象")
	})
	return instance
}

// 示例方法
func (s *Singleton) DoSomething() {
	fmt.Println("单例对象在执行任务")
}
