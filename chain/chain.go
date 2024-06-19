package chain

import "fmt"

// Handler 处理器接口
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler 基础处理器结构体
type BaseHandler struct {
	nextHandler Handler
}

// SetNext 设置下一个处理器
func (h *BaseHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

// Handle 处理请求
func (h *BaseHandler) Handle(request string) {
	if h.nextHandler != nil {
		h.nextHandler.Handle(request)
	}
}

// ConcreteHandler1 具体处理器1
type ConcreteHandler1 struct {
	BaseHandler
}

// Handle 处理请求
func (h *ConcreteHandler1) Handle(request string) {
	if request == "request1" {
		fmt.Println("ConcreteHandler1 handled the request")
		return
	}
	fmt.Println("ConcreteHandler1 passed the request")
	h.BaseHandler.Handle(request)
}

// ConcreteHandler2 具体处理器2
type ConcreteHandler2 struct {
	BaseHandler
}

// Handle 处理请求
func (h *ConcreteHandler2) Handle(request string) {
	if request == "request2" {
		fmt.Println("ConcreteHandler2 handled the request")
		return
	}
	fmt.Println("ConcreteHandler2 passed the request")
	h.BaseHandler.Handle(request)
}
