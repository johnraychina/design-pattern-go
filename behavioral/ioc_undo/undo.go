package main

import (
	"errors"
	"fmt"
)

func main() {
	set := NewIntSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Undo()
	set.Undo()
	fmt.Printf("%#v", set)
}

// type Undo: a function list
// method Add: add function to the list
// here we abstract Undo as a type, so that it depends on nothing specific biz logic
// and conversely, biz logic depends on Undo type.

type Undo []func()

func (undo *Undo) Add(f func()) {
	*undo = append(*undo, f)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("no function to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // for garbage collection
	}
	*undo = functions[:index]
	return nil
}

type IntSet struct {
	data map[int]bool
	undo Undo
}

func NewIntSet() IntSet {
	// todo slice 不用make吗？
	return IntSet{data: make(map[int]bool)}
}

func (i *IntSet) Undo() {
	i.undo.Undo()
}

func (i *IntSet) contains(x int) bool {
	return i.data[x]
}

func (i *IntSet) Add(x int) {
	if !i.contains(x) {
		i.data[x] = true
		i.undo.Add(func() { i.Delete(x) })
	} else {
		i.undo.Add(nil)
	}
}

func (i *IntSet) Delete(x int) {
	if i.contains(x) {
		delete(i.data, x)
		i.undo.Add(func() { i.Add(x) })
	} else {
		i.undo.Add(nil)
	}
}
