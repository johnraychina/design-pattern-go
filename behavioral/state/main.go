package main

// 为了取代复杂的if状态转换，抽象出来状态。
// 每个状态包含可推进的下一组状态，每个状态只关心自己遇到某个事件如何流转到下个状态。
func main() {

}

type Machine struct {
	currentState state
}

type state interface {
	onEvent(eventType int)
}

type Ready interface {
	onEvent(eventType int)
}
type Stopped interface {
	onEvent(eventType int)
}
type Running interface {
	onEvent(eventType int)
}
