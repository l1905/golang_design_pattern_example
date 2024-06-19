package template

import "fmt"

// 定义一个抽象的模板接口
type Template interface {
	Step1()
	Step2()
	Step3()
}

// 实现一个具体的模板结构体
type ConcreteTemplate struct{}

func (t *ConcreteTemplate) Step1() {
	fmt.Println("执行步骤 1")
}

func (t *ConcreteTemplate) Step2() {
	fmt.Println("执行步骤 2")
}

func (t *ConcreteTemplate) Step3() {
	fmt.Println("执行步骤 3")
}

// 定义一个算法的模板方法
func TemplateMethod(t Template) {
	t.Step1()
	t.Step2()
	t.Step3()
}

// 定义具体的实现类
type ConcreteImplementationA struct {
	ConcreteTemplate
}

// 实现具体的步骤2
func (a *ConcreteImplementationA) Step2() {
	fmt.Println("ConcreteImplementationA 执行步骤 2")
}

type ConcreteImplementationB struct {
	ConcreteTemplate
}

// 显现具体的步骤3
func (b *ConcreteImplementationB) Step3() {
	fmt.Println("ConcreteImplementationB 执行步骤 3")
}
