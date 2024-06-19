package memento

import "fmt"

// Memento 存储对象的状态
type Memento struct {
	state string
}

// NewMemento 创建一个新的备忘录
func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

// GetState 返回备忘录存储的状态
func (m *Memento) GetState() string {
	return m.state
}

// Originator 创建状态并生成备忘录
type Originator struct {
	state string
}

// SetState 设置当前状态
func (o *Originator) SetState(state string) {
	o.state = state
	fmt.Printf("Setting state to: %s\n", state)
}

// SaveStateToMemento 保存当前状态到备忘录
func (o *Originator) SaveStateToMemento() *Memento {
	return NewMemento(o.state)
}

// GetStateFromMemento 从备忘录恢复状态
func (o *Originator) GetStateFromMemento(memento *Memento) {
	o.state = memento.GetState()
	fmt.Printf("Restoring state to: %s\n", o.state)
}

// Caretaker 负责保存和恢复多个备忘录
type Caretaker struct {
	mementoList []*Memento
}

// Add 添加备忘录到列表
func (c *Caretaker) Add(m *Memento) {
	c.mementoList = append(c.mementoList, m)
}

// Get 获取特定索引的备忘录
func (c *Caretaker) Get(index int) *Memento {
	return c.mementoList[index]
}
