package main

import "fmt"

// 避免receiver与invoker多对多，在receiver和invoker中间加一个command层，
// 这样你可以组合invoker-command，不需要每个invoker都实现command的逻辑，减少invoker的冗余
// 真实世界中，在大的餐厅，不会每个顾客直接对厨师提需求，你去餐厅吃饭，通常是先下单，然后订单被排队，厨师按订单做好后，服务员再送上桌。
// 这个订单就是你的command，一个接口产生不同行为，command可以被delay或者排队执行，以及支持撤销的操作。
// Command is a behavioral design pattern that turns a request into a stand-alone object
// that contains all information about the request.
// This transformation lets you pass requests as a method arguments, delay or queue a request’s execution,
// and support undoable operations.
// 适用场景：
//  Use the Command pattern when you want to parametrize objects with operations.
//  Use the Command pattern when you want to queue operations, schedule their execution, or execute them remotely.
//  Use the Command pattern when you want to implement reversible operations.

func main() {
	t := &TV{}
	on := &onCommand{device: t}
	off := &offCommand{device: t}
	onBtn := &Button{command: on}
	offBtn := &Button{command: off}
	onBtn.Press()
	offBtn.Press()
}

// Device : Receiver interface
type Device interface {
	On()
	Off()
}

// TV : Concrete receiver
type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}
func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// Command : Command interface
type Command interface {
	execute()
}

// onCommand : Concrete command
type onCommand struct {
	device Device
}

func (c *onCommand) execute() {
	c.device.On()
}

type offCommand struct {
	device Device
}

func (c *offCommand) execute() {
	c.device.Off()
}

// Button : Invoker
type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.execute()
}
