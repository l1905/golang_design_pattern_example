package agent

// Subject 定义了 RealSubject 和 Proxy 的共用接口,这样就在任何使用 RealSubject 的地方都可以使用 Proxy
type Subject interface {
	Do() string
}

// RealSubject 定义 Proxy 所代表的真实实体
type RealSubject struct{}

func (RealSubject) Do() string {
	return "Real Subject"
}

// Proxy 保存一个引用使得代理可以访问实体,并提供一个与 Subject 的接口相同的接口,这样代理就可以用来代替实体
type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前可以做一些额外的事情

	res += "Proxy: Before:"

	// 调用真实对象的方法
	res += p.real.Do()

	// 调用真实对象之后可以做一些额外的事情
	res += ":After"

	return res
}
