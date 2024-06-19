package iterator

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// ConcreteIterator 具体迭代器
type ConcreteIterator struct {
	data  []interface{}
	index int
}

// HasNext 判断是否还有下一个元素
func (iter *ConcreteIterator) HasNext() bool {
	return iter.index < len(iter.data)
}

// Next 返回当前元素并移动到下一个位置
func (iter *ConcreteIterator) Next() interface{} {
	if iter.HasNext() {
		item := iter.data[iter.index]
		iter.index++
		return item
	}
	return nil
}

// Aggregate 聚合对象接口
type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteAggregate 具体聚合对象
type ConcreteAggregate struct {
	Data []interface{}
}

// CreateIterator 创建迭代器
func (agg *ConcreteAggregate) CreateIterator() Iterator {
	return &ConcreteIterator{data: agg.Data}
}
