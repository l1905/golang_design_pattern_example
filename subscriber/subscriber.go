package subscriber

import "fmt"

// 定义订阅者接口
type Subscriber interface {
	Update(message string)
}

// 定义具体的订阅者类型
type ConcreteSubscriber struct {
	Name string
}

func (s *ConcreteSubscriber) Update(message string) {
	fmt.Printf("%s received message: %s\n", s.Name, message)
}

// 定义发布者类型
type Publisher struct {
	subscribers []Subscriber
}

func (p *Publisher) AddSubscriber(subscriber Subscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p *Publisher) RemoveSubscriber(subscriber Subscriber) {
	for i, s := range p.subscribers {
		if s == subscriber {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
			break
		}
	}
}

func (p *Publisher) Notify(message string) {
	for _, subscriber := range p.subscribers {
		subscriber.Update(message)
	}
}
