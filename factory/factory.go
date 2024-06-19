package factory

import "fmt"

// 创建型模式-工厂模式
// 定义一个接口
type Shape interface {
	Draw()
}

// 定义具体的产品类型
type Circle struct{}
type Square struct{}

// 实现 Shape 接口的 Draw 方法
func (c Circle) Draw() {
	fmt.Println("画一个圆形")
}

func (s Square) Draw() {
	fmt.Println("画一个正方形")
}

// 简单工厂函数
func ShapeSimpleFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "square":
		return Square{}
	default:
		return nil
	}
}

// 定义常规工厂接口
type ShapeFactory interface {
	CreateShape() Shape
}

// 常规的工厂函数
type CircleFactory struct{}
type SquareFactory struct{}

// 实现 ShapeFactory 接口的 CreateShape 方法
func (cf CircleFactory) CreateShape() Shape {
	return Circle{}
}

func (sf SquareFactory) CreateShape() Shape {
	return Square{}
}
