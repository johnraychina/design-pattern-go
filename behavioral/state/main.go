package main

import (
	"fmt"
	"strconv"
)

// 为了取代复杂的if状态转换，抽象出来状态。
// 每个状态包含可推进的下一组状态，每个状态只关心自己遇到某个事件如何流转到下个状态。
// 状态转换可以简单的if-else映射到下个状态，也可以通过map映射到下个状态
// 后者更加灵活，可以做成配置化。
// 最后甚至可以就是一个 map[当前状态]map[事件类型]下个状态

func main() {

	m := &Machine{}
	s1 := &Ready{m: m}
	s2 := &Running{m: m}
	s3 := &Stopped{m: m}

	m.currentState = s3
	m.ready = s1
	m.running = s2
	m.stopped = s3

	m.Start()
	m.Run()
	m.Stop()

	m.Run()
}

type Machine struct {
	// 将所有状态枚举出来
	ready   State
	running State
	stopped State

	currentState State
	execCount    int
}

func (m *Machine) Start() {
	fmt.Println("start")
	m.currentState.onEvent(EventTypeStart)
}
func (m *Machine) Run() {
	fmt.Println("run")
	m.currentState.onEvent(EventTypeRun)
}
func (m *Machine) Stop() {
	fmt.Println("stop")
	m.currentState.onEvent(EventTypeStop)
}

const (
	EventTypeStart = 1
	EventTypeRun   = 2
	EventTypeStop  = 3
)

type State interface {
	onEvent(eventType int)
}

type Ready struct {
	m *Machine
}

func (s *Ready) onEvent(eventType int) {

	if eventType == EventTypeStart {
		// 继续ready
	} else if eventType == EventTypeRun {
		s.m.currentState = s.m.running
	} else if eventType == EventTypeStop {
		s.m.currentState = s.m.stopped
	} else {
		panic("invalid event type for ready state:" + strconv.Itoa(eventType))
	}
}

type Running struct {
	m *Machine
}

func (s *Running) onEvent(eventType int) {
	if eventType == EventTypeRun {
		// 继续run
	} else if eventType == EventTypeStop {
		s.m.currentState = s.m.stopped
	} else {
		panic("invalid event type for ready state:" + strconv.Itoa(eventType))
	}
}

type Stopped struct {
	m *Machine
}

func (s *Stopped) onEvent(eventType int) {
	if eventType == EventTypeStop {
		// 继续stopped
	} else if eventType == EventTypeStart {
		s.m.currentState = s.m.running
	} else {
		panic("invalid event type for stopped state:" + strconv.Itoa(eventType))
	}
}
