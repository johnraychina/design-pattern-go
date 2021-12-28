package main

import "fmt"

// Chain of responsibility is a behavioral design pattern that lets you
// pass requests along a chain of handlers. Upon receiving request,
// each handler decides either process request or to pass it to the next handler in the chain.
func main() {
	h3 := &ThirdHandler{}
	h2 := &SecondHandler{Next: h3}
	h1 := &FirstHandler{Next: h2}
	h1.Handle(&Request{})
}

type Request struct {
	name   string
	first  bool
	second bool
	third  bool
}

type Handler interface {
	SetNext(handler Handler)
	Handle(*Request)
}

type FirstHandler struct {
	Next Handler
}

func (h *FirstHandler) SetNext(next Handler) {
	h.Next = next
}

func (h *FirstHandler) Handle(request *Request) {
	if request.first == true {
		h.Next.Handle(request)
		return
	}
	fmt.Println("FirstHandler handling request")
	request.first = true
	h.Next.Handle(request)
	return
}

type SecondHandler struct {
	Next Handler
}

func (h *SecondHandler) SetNext(next Handler) {
	h.Next = next
}

func (h *SecondHandler) Handle(request *Request) {
	if request.second == true {
		h.Next.Handle(request)
		return
	}
	fmt.Println("SecondHandler handling request")
	request.second = true
	h.Next.Handle(request)
	return
}

type ThirdHandler struct {
	Next Handler
}

func (h *ThirdHandler) SetNext(next Handler) {
	h.Next = next
}

func (h *ThirdHandler) Handle(request *Request) {
	if request.third == true {
		return
	}
	fmt.Println("ThirdHandler handling request")
	request.third = true
	return
}
