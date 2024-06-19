package main

import (
	"design_pattern/adapter"
	"design_pattern/agent"
	"design_pattern/bridge"
	"design_pattern/builder"
	"design_pattern/chain"
	"design_pattern/command"
	"design_pattern/composite"
	"design_pattern/decorator"
	"design_pattern/facade"
	"design_pattern/factory"
	"design_pattern/flyweight"
	"design_pattern/interpreter"
	"design_pattern/iterator"
	"design_pattern/memento"
	"design_pattern/prototype"
	"design_pattern/singleton"
	"design_pattern/state"
	"design_pattern/strategy"
	"design_pattern/subscriber"
	"design_pattern/template"
	"design_pattern/visitor"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {

	// builderDemo() // 构造者模式
	// prototypeDemo() // 原型模式
	// adapterDemo() // 适配器模式
	// agentDemo() // 代理模式
	// decoratorDemo() // 装饰器模式

	// facadeDemo() // 门面模式

	// bridgeDemo() // 桥接模式

	// compositeDemo() // 组合模式

	// flyweightDemo() // 享元模式

	// strategyDemo() // 策略模式

	// subscriberDemo() // 订阅者模式

	// chainOfResponsibilityDemo() // 责任链模式

	// commandDemo() // 命令模式

	// iteratorDemo() // 迭代器模式

	// templateDemo() // 模板模式

	// stateDemo() // 状态模式

	// visitorDemo() // 访问者模式

	// visitorAstDemo() // 访问者模式-ast

	// mementoDemo() // 备忘录模式

	interpreterDemo() // 解释器模式
}

// 单例模式
func singletonDemo() {
	singletonInstance := singleton.GetInstance()
	singletonInstance.DoSomething()

	singletonInstance2 := singleton.GetInstance()
	singletonInstance2.DoSomething()
}

// 简单工厂模式
func simpleFactoryDemo() {

	// 简单工厂模式
	circle := factory.ShapeSimpleFactory("circle")
	circle.Draw()

	square := factory.ShapeSimpleFactory("square")
	square.Draw()
}

// 工厂模式
func factoryDemo() {
	var ShapeFactory factory.ShapeFactory
	// 创建 ShapeFactory 对象
	ShapeFactory = factory.CircleFactory{}
	// 使用 ShapeFactory 创建 Circle 对象
	circle := ShapeFactory.CreateShape()
	circle.Draw()

	// 创建 ShapeFactory 对象
	ShapeFactory = factory.SquareFactory{}
	// 使用 SquareFactory 创建 Square 对象
	square := ShapeFactory.CreateShape()
	square.Draw()
}

// 建造者模式
func builderDemo() {
	// 创建指挥者对象
	director := builder.Director{}
	// 创建具体的建造者对象
	builder := &builder.ConcreteBuilder{}
	// 指挥者使用建造者对象来构建产品
	director.SetBuilder(builder)
	product := director.Construct()
	// 使用产品
	fmt.Println(product)
}

// 原型模式
func prototypeDemo() {
	// 创建一个具体的原型对象
	prototype := &prototype.ConcretePrototype{
		Name: "prototype",
	}
	// 克隆原型对象
	clone := prototype.Clone()
	fmt.Println(clone)
}

// adapter适配器模式
func adapterDemo() {
	adaptee := &adapter.Adaptee{}
	adapterInstance := &adapter.Adapter{Adaptee: adaptee}

	adapter.Client(adapterInstance)
}

// 代理模式
func agentDemo() {
	var sub agent.Subject

	sub = &agent.RealSubject{}
	res := sub.Do()
	fmt.Println(res)

	// 代理模式
	sub = &agent.Proxy{}
	res = sub.Do()
	fmt.Println(res)
}

// 装饰器模式
func decoratorDemo() {
	// 创建一个具体的组件
	component := &decorator.ConcreteComponent{}

	// 使用装饰器 A 装饰组件
	decoratorA := &decorator.ConcreteDecoratorA{decorator.Decorator{component}}

	// 使用装饰器 B 装饰已经被装饰器 A 装饰过的组件
	decoratorB := &decorator.ConcreteDecoratorB{decorator.Decorator{decoratorA}}

	// 调用最外层装饰器的方法
	result := decoratorB.Operation()
	fmt.Println(result)
}

// 外观设计模式
func facadeDemo() {
	// 创建一个外观对象
	facade := facade.NewFacade()

	// 调用外观对象的方法
	facade.OperationWrapper()
}

// 桥接模式
func bridgeDemo() {
	redColor := &bridge.RedColor{}
	blueColor := &bridge.BlueColor{}

	circleRed := &bridge.Circle{Color: redColor}
	circleBlue := &bridge.Circle{Color: blueColor}
	squareRed := &bridge.Square{Color: redColor}
	squareBlue := &bridge.Square{Color: blueColor}

	circleRed.Draw()
	circleBlue.Draw()
	squareRed.Draw()
	squareBlue.Draw()
}

// 组合模式
func compositeDemo() {
	root := &composite.Composite{Name: "Root"}

	branch1 := &composite.Composite{Name: "Branch1"}
	branch1.Add(&composite.Leaf{Name: "Leaf1"})
	branch1.Add(&composite.Leaf{Name: "Leaf2"})

	branch2 := &composite.Composite{Name: "Branch2"}
	branch2.Add(&composite.Leaf{Name: "Leaf3"})

	root.Add(branch1)
	root.Add(branch2)

	root.Operation()
}

// 享元模式-数据共享
func flyweightDemo() {
	factory := flyweight.NewShapeFactory()

	// 获取红色圆形并绘制
	redCircle := factory.GetCircle("red")
	redCircle.Draw()

	// 获取蓝色圆形并绘制
	blueCircle := factory.GetCircle("blue")
	blueCircle.Draw()

	// 再次获取红色圆形并绘制
	anotherRedCircle := factory.GetCircle("red")
	anotherRedCircle.Draw()
}

// 行为设计模式-策略模式
func strategyDemo() {
	context := &strategy.Context{}

	// 使用add策略
	add := &strategy.Add{}
	context.SetStrategy(add)
	result := context.ExecuteStrategy(5, 3)
	fmt.Printf("Add: %d\n", result)

	subtract := &strategy.Subtract{}
	// 使用subtract策略
	context.SetStrategy(subtract)
	result = context.ExecuteStrategy(5, 3)
	fmt.Printf("Subtract: %d\n", result)

	multiply := &strategy.Multiply{}
	// 使用multiply策略
	context.SetStrategy(multiply)
	result = context.ExecuteStrategy(5, 3)
	fmt.Printf("Multiply: %d\n", result)
}

// 订阅者模式
func subscriberDemo() {
	// 创建发布者
	publisher := &subscriber.Publisher{}

	// 创建订阅者
	subscriber1 := &subscriber.ConcreteSubscriber{Name: "Subscriber 1"}
	subscriber2 := &subscriber.ConcreteSubscriber{Name: "Subscriber 2"}

	// 订阅者订阅发布者
	publisher.AddSubscriber(subscriber1)
	publisher.AddSubscriber(subscriber2)

	// 发布者发布消息
	publisher.Notify("Hello, subscribers!")

	// 移除订阅者
	publisher.RemoveSubscriber(subscriber2)

	// 再次发布消息
	publisher.Notify("Another message")
}

// 责任链模式
func chainOfResponsibilityDemo() {
	handler1 := &chain.ConcreteHandler1{}
	handler2 := &chain.ConcreteHandler2{}

	handler1.SetNext(handler2)

	requests := []string{"request1", "request2", "request3"}
	for _, request := range requests {
		handler1.Handle(request)
		fmt.Println()
	}
}

// 命令模式
func commandDemo() {
	calculator := &command.Calculator{}

	addCommand := &command.AddCommand{Calculator: calculator, Operand: 10}
	subtractCommand := &command.SubtractCommand{Calculator: calculator, Operand: 5}

	invoker := &command.Invoker{}
	invoker.AddCommand(addCommand)
	invoker.AddCommand(subtractCommand)

	invoker.ExecuteCommands()

	fmt.Println("Result:", calculator.GetResult())
}

// 迭代器模式
func iteratorDemo() {
	// 创建聚合对象
	aggregate := &iterator.ConcreteAggregate{
		Data: []interface{}{1, 2, 3, 4, 5},
	}

	// 创建迭代器
	iterator := aggregate.CreateIterator()

	// 遍历聚合对象
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}

// 模版设计模式
func templateDemo() {
	implementationA := &template.ConcreteImplementationA{}
	template.TemplateMethod(implementationA)

	fmt.Println("------------------------")

	implementationB := &template.ConcreteImplementationB{}
	template.TemplateMethod(implementationB)
}

// 状态设计模式
func stateDemo() {
	order := &state.Order{State: &state.PendingState{}}
	fmt.Printf("当前订单状态: %s\n", order.GetState())

	order.Pay()
	order.SetState(&state.PaidState{})
	fmt.Printf("当前订单状态: %s\n", order.GetState())

	order.Ship()
	order.SetState(&state.ShippedState{})
	fmt.Printf("当前订单状态: %s\n", order.GetState())

	order.Complete()
	order.SetState(&state.CompletedState{})
	fmt.Printf("当前订单状态: %s\n", order.GetState())
}

// 访问者模式
func visitorDemo() {
	elements := []visitor.Element{
		&visitor.ConcreteElementA{Name: "Element A1"},
		&visitor.ConcreteElementB{Value: 10},
		&visitor.ConcreteElementA{Name: "Element A2"},
	}

	visitor := &visitor.ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}

// 访问者模式-ast
func visitorAstDemo() {
	src := `
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, world!")
    }

    func add(a int, b int) int {
        return a + b
    }
    `

	// 创建文件集
	fset := token.NewFileSet()

	// 解析源代码
	f, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建并使用访问者遍历语法树
	v := &visitor.FuncVisitor{}
	ast.Walk(v, f)
}

// 备忘录模式
func mementoDemo() {
	originator := &memento.Originator{}
	caretaker := &memento.Caretaker{}

	// 设置状态并保存到备忘录
	originator.SetState("State #1")
	caretaker.Add(originator.SaveStateToMemento())

	originator.SetState("State #2")
	caretaker.Add(originator.SaveStateToMemento())

	originator.SetState("State #3")
	caretaker.Add(originator.SaveStateToMemento())

	// 恢复状态
	originator.GetStateFromMemento(caretaker.Get(0))
	originator.GetStateFromMemento(caretaker.Get(1))
}

// 解释器模式
func interpreterDemo() {
	template := "Hello, {{name}} ! Welcome to {{site}} ."
	context := &interpreter.Context{
		Variables: map[string]string{
			"name": "Alice",
			"site": "GoLand",
		},
	}

	expressions := interpreter.Template(template)
	d, _ := json.Marshal(expressions)
	fmt.Println(string(d))
	result := interpreter.Render(expressions, context)
	fmt.Println(result)
}
