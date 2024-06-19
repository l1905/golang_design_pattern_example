package bridge

import "fmt"

// 颜色接口
type Color interface {
	applyColor()
}

// 红色
type RedColor struct{}

func (r *RedColor) applyColor() {
	fmt.Println("Applying red color")
}

// 蓝色
type BlueColor struct{}

func (b *BlueColor) applyColor() {
	fmt.Println("Applying blue color")
}

// 形状接口
type Shape interface {
	draw()
}

// 圆形
type Circle struct {
	Color Color
}

func (c *Circle) Draw() {
	fmt.Println("Drawing circle")
	c.Color.applyColor()
}

// 正方形
type Square struct {
	Color Color
}

func (s *Square) Draw() {
	fmt.Println("Drawing square")
	s.Color.applyColor()
}
