package facade

import "fmt"

// 结构设计模式-外观设计模式 Subsystem1 是子系统1
type Subsystem1 struct{}

func (s *Subsystem1) Operation1() {
	fmt.Println("Subsystem1: Ready!")
	fmt.Println("Subsystem1: Go!")
}

// Subsystem2 是子系统2
type Subsystem2 struct{}

func (s *Subsystem2) Operation2() {
	fmt.Println("Subsystem2: Get ready!")
	fmt.Println("Subsystem2: Fire!")
}

// Facade 是外观类
type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

// NewFacade 创建一个新的外观类实例
func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

// OperationWrapper 封装子系统操作的高层接口
func (f *Facade) OperationWrapper() {
	fmt.Println("Facade initializes subsystems:")
	f.subsystem1.Operation1()
	f.subsystem2.Operation2()
	fmt.Println("Facade orders subsystems to perform the action:")
	f.subsystem1.Operation1()
	f.subsystem2.Operation2()
}
