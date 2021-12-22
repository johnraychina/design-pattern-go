package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range sum(sq(odd(echo(nums)))) {
		fmt.Println(n)
	}
}

// echo: input(datasource), output channel
func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq: input channel, square int, output channel
func sq(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// odd: input channel, filter odd, output channel
func odd(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			if n&1 == 1 {
				out <- n * n
			}
		}
		close(out)
	}()
	return out
}

// sum
func sum(ch <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		var sum = 0
		for n := range ch {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
