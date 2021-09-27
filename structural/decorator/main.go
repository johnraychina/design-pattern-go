package main

import "fmt"

func main() {
	pizza := &veggeMania{}

	// Add cheese topping
	// 注意，这里要取地址&，才会成为interface
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	// Add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
