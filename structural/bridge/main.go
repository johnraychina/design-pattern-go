package main

import "fmt"

// Bridge is a structural design pattern that lets you split a large class
//or a set of closely related classes into two separate hierarchies—abstraction
//and implementation—which can be developed independently of each other.

// 解决组合爆炸和独立演化问题
// 例子：shape + color ==> RedCircle, RedSquare, BlueCircle, BlueSquare
// 遥控器 --> 设备
// GUI风格 --> 操作系统
// computer --> printer

func main() {

	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}

	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	winComputer := &windows{}

	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}


