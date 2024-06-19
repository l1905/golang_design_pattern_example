package adapter

import "fmt"

// Target 定义客户端期望的接口
type Target interface {
	Request() string
}

// Adaptee 是需要适配的类
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee's specific request"
}

// Adapter 是适配器，实现了 Target 接口
type Adapter struct {
	Adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

// Client 是客户端，使用 Target 接口
func Client(target Target) {
	fmt.Println(target.Request())
}
