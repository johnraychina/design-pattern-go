package main

import (
	"fmt"
)

// a visitor function
// a visitor interface
// a struct(Info) that implements the visitor interface
// visitors that do different things

type VisitorFunc func(*Info) error

type Visitor interface {
	Visit(visitorFunc VisitorFunc) error
}

type Info struct {
	Namespace string
	Name      string
	Other     string
}

func (i *Info) Visit(visitorFunc VisitorFunc) error {
	return visitorFunc(i)
}

// NameVisitor 使用功能decorator 模式：包住下一层visitor，自己也实现Visitor接口
type NameVisitor struct {
	visitor Visitor
}

func (v *NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info) error {
		fmt.Println("Name:", info.Name)
		fmt.Println("NameVisitor() before calling function")
		err := fn(info)
		fmt.Println("NameVisitor() after calling function")
		fmt.Println("Name:", info.Name)
		return err
	})
}

type LogVisitor struct {
	visitor Visitor
}

func (v *LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info) error {
		fmt.Println("LogVisitor() before calling function")
		err := fn(info)
		fmt.Println("LogVisitor() after calling function")
		return err
	})
}

func main() {
	info := &Info{}

	var v Visitor = &LogVisitor{info}
	v = &NameVisitor{v}
	v.Visit(func(i *Info) error {
		i.Namespace = "github"
		i.Name = "johnraychina"
		i.Other = "Doing research"
		fmt.Println("ok, info is loaded!")
		return nil
	})
}
