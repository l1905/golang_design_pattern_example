package visitor

import "fmt"

// Element 接口定义了一个接受访问者的方法
type Element interface {
	Accept(Visitor)
}

// Visitor 接口定义了代表访问者的所有操作
type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

// ConcreteElementA 是一个具体元素
type ConcreteElementA struct {
	Name string
}

// Accept 方法实现了 Element 接口，接受一个访问者
func (e *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(e)
}

// ConcreteElementB 是另一个具体元素
type ConcreteElementB struct {
	Value int
}

// Accept 方法实现了 Element 接口，接受一个访问者
func (e *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(e)
}

// ConcreteVisitor 是一个具体访问者
type ConcreteVisitor struct{}

// VisitConcreteElementA 实现了访问者对 ConcreteElementA 的操作
func (v *ConcreteVisitor) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Printf("Visiting ConcreteElementA: %s\n", e.Name)
}

// VisitConcreteElementB 实现了访问者对 ConcreteElementB 的操作
func (v *ConcreteVisitor) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Printf("Visiting ConcreteElementB: %d\n", e.Value)
}
