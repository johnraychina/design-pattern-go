package main

import "fmt"

// Flyweight is a structural design pattern that allows programs
// to support vast quantities of objects by keeping their memory
// consumption low.

// The pattern achieves it
// by sharing parts of object state between multiple objects.
// In other words, the Flyweight saves RAM
// by caching the same data used by different objects.
// 减小内存消耗，让大量对象共享相同部分的内存.
// 例子：CS游戏，玩家的皮肤，可以共享
func main() {

	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
