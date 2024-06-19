package command

// 命令的接口形式
type Command interface {
	Execute()
}

// 加操作命令
type AddCommand struct {
	Calculator *Calculator // 依赖命令的接收者
	Operand    int
}

// 实现命令接口
func (c *AddCommand) Execute() {
	c.Calculator.Add(c.Operand)
}

// 减操作命令
type SubtractCommand struct {
	Calculator *Calculator
	Operand    int
}

func (c *SubtractCommand) Execute() {
	c.Calculator.Subtract(c.Operand)
}

// 计算器结构体
type Calculator struct {
	result int
}

func (c *Calculator) Add(operand int) {
	c.result += operand
}

func (c *Calculator) Subtract(operand int) {
	c.result -= operand
}

func (c *Calculator) GetResult() int {
	return c.result
}

// 调用者
type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}
