package strategy

// Strategy 接口定义了策略的行为
type Strategy interface {
	Execute(int, int) int
}

// Add 结构体实现了 Strategy 接口，表示加法策略
type Add struct{}

func (a *Add) Execute(x, y int) int {
	return x + y
}

// Subtract 结构体实现了 Strategy 接口，表示减法策略
type Subtract struct{}

func (s *Subtract) Execute(x, y int) int {
	return x - y
}

// Multiply 结构体实现了 Strategy 接口，表示乘法策略
type Multiply struct{}

func (m *Multiply) Execute(x, y int) int {
	return x * y
}

// Context 结构体持有一个 Strategy 接口，可以动态地改变策略
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(x, y int) int {
	return c.strategy.Execute(x, y)
}
