package composite

import "fmt"

// 结构型设计模式-组合模式

// 核心是类似二叉树、或者树形结构的数据结构，通过递归的方式构建整个对象，使得整个对象的使用方式一致

// Component 接口定义了组合对象和叶子对象的共同操作
type Component interface {
	Operation()
}

// Leaf 结构体表示叶子对象,实现了 Component 接口
type Leaf struct {
	Name string
}

func (l *Leaf) Operation() {
	fmt.Printf("Leaf %s: Operation\n", l.Name)
}

// Composite 结构体表示组合对象,实现了 Component 接口
type Composite struct {
	Name     string
	Children []Component
}

func (c *Composite) Operation() {
	fmt.Printf("Composite %s: Operation\n", c.Name)
	for _, child := range c.Children {
		child.Operation()
	}
}

func (c *Composite) Add(component Component) {
	c.Children = append(c.Children, component)
}

func (c *Composite) Remove(component Component) {
	for i, child := range c.Children {
		if child == component {
			c.Children = append(c.Children[:i], c.Children[i+1:]...)
			return
		}
	}
}
