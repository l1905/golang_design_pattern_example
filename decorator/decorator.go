package decorator

type Component interface {
	Operation() string
}

// ConcreteComponent 是一个具体的组件实现
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// Decorator 是一个装饰器，它包含一个 Component 引用
type Decorator struct {
	Component Component
}

func (d *Decorator) Operation() string {
	return d.Component.Operation()
}

// ConcreteDecoratorA 是一个具体的装饰器实现
type ConcreteDecoratorA struct {
	Decorator
}

func (d *ConcreteDecoratorA) Operation() string {
	return "ConcreteDecoratorA(" + d.Component.Operation() + ")"
}

// ConcreteDecoratorB 是另一个具体的装饰器实现
type ConcreteDecoratorB struct {
	Decorator
}

func (d *ConcreteDecoratorB) Operation() string {
	return "ConcreteDecoratorB(" + d.Component.Operation() + ")"
}
