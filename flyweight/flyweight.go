package flyweight

import "fmt"

// 享元接口
type Shape interface {
	Draw()
}

// 具体享元结构体
type Circle struct {
	color string
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a %s circle\n", c.color)
}

// 享元工厂结构体
type ShapeFactory struct {
	circleMap map[string]*Circle
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		circleMap: make(map[string]*Circle),
	}
}

func (f *ShapeFactory) GetCircle(color string) *Circle {
	if circle, ok := f.circleMap[color]; ok {
		return circle
	}
	circle := &Circle{color: color}
	f.circleMap[color] = circle
	fmt.Printf("Creating a %s circle\n", color)
	return circle
}
