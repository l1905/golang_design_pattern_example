package builder

// 构造者模式

// Product 代表要构建的复杂对象
type Product struct {
	part1 string
	part2 string
	part3 string
}

// Builder 是建造者接口
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetProduct() Product
}

// ConcreteBuilder 是具体的建造者实现
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPart1() {
	b.product.part1 = "Part1"
}

func (b *ConcreteBuilder) BuildPart2() {
	b.product.part2 = "Part2"
}

func (b *ConcreteBuilder) BuildPart3() {
	b.product.part3 = "Part3"
}

func (b *ConcreteBuilder) GetProduct() Product {
	return b.product
}

// Director 是指挥者,负责控制建造过程
type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Construct() Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
	return d.builder.GetProduct()
}
