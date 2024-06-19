package prototype

// 创建型模式-原型模式

// Cloneable 是原型接口
type Cloneable interface {
	Clone() Cloneable
}

// ConcretePrototype 是具体的原型结构体
type ConcretePrototype struct {
	Name string
}

// Clone 方法实现了 Cloneable 接口,用于克隆一个新的对象
func (p *ConcretePrototype) Clone() Cloneable {
	return &ConcretePrototype{
		Name: p.Name + "_clone",
	}
}
