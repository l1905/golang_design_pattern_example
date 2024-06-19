package state

import "fmt"

// 状态模式 OrderState 订单状态接口
type OrderState interface {
	Pay()
	Ship()
	Complete()
	GetState() string
}

// Order 订单结构体
type Order struct {
	State OrderState
}

func (o *Order) SetState(state OrderState) {
	o.State = state
}

func (o *Order) Pay() {
	o.State.Pay()
}

func (o *Order) Ship() {
	o.State.Ship()
}

func (o *Order) Complete() {
	o.State.Complete()
}

func (o *Order) GetState() string {
	return o.State.GetState()
}

// PendingState 待付款状态
type PendingState struct{}

func (s *PendingState) Pay() {
	fmt.Println("订单已付款,状态变为已付款")
}

func (s *PendingState) Ship() {
	fmt.Println("订单未付款,无法发货")
}

func (s *PendingState) Complete() {
	fmt.Println("订单未发货,无法完成")
}

func (s *PendingState) GetState() string {
	return "待付款"
}

// PaidState 已付款状态
type PaidState struct{}

func (s *PaidState) Pay() {
	fmt.Println("订单已经付款,无需重复付款")
}

func (s *PaidState) Ship() {
	fmt.Println("订单已发货,状态变为发货中")
}

func (s *PaidState) Complete() {
	fmt.Println("订单未发货,无法完成")
}

func (s *PaidState) GetState() string {
	return "已付款"
}

// ShippedState 发货中状态
type ShippedState struct{}

func (s *ShippedState) Pay() {
	fmt.Println("订单已经付款,无需重复付款")
}

func (s *ShippedState) Ship() {
	fmt.Println("订单已经发货,无需重复发货")
}

func (s *ShippedState) Complete() {
	fmt.Println("订单已完成,状态变为已完成")
}

func (s *ShippedState) GetState() string {
	return "发货中"
}

// CompletedState 已完成状态
type CompletedState struct{}

func (s *CompletedState) Pay() {
	fmt.Println("订单已经完成,无需付款")
}

func (s *CompletedState) Ship() {
	fmt.Println("订单已经完成,无需发货")
}

func (s *CompletedState) Complete() {
	fmt.Println("订单已经完成,无需重复完成")
}

func (s *CompletedState) GetState() string {
	return "已完成"
}
