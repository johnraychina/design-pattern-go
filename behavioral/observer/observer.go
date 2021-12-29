package main

import "fmt"

// 好莱坞原则：don't call me, I'll call you
// 避免大量频繁轮询，而是先注册，后通知
func main() {
	publisher := &DeliverCompany{}
	supplier := &Supplier{}
	customer := &Customer{}

	publisher.Subscribe(supplier)
	publisher.Subscribe(customer)

	publisher.NotifyAll(Order{id: "x001", name: "iPhone 13"})
}

type Publisher interface {
	// 这里可以订阅特定事件
	//subscribe(eventType string, Subscriber)

	Subscribe(Subscriber)
	NotifyAll(Order)
}

type Subscriber interface {
	Notify(Order)
}

type Order struct {
	id   string
	name string
}

type DeliverCompany struct {
	subscribers []Subscriber
}

func (d *DeliverCompany) Subscribe(s Subscriber) {
	d.subscribers = append(d.subscribers, s)
}
func (d *DeliverCompany) NotifyAll(msg Order) {
	for _, s := range d.subscribers {
		s.Notify(msg)
	}
}

type Customer struct {
}

func (c *Customer) Notify(message Order) {
	fmt.Println("ok, I'll pickup the package after 6:00pm")
}

type Supplier struct {
}

func (s *Supplier) Notify(message Order) {
	fmt.Println("ok, I'll tag this order as [delivered]")
}
