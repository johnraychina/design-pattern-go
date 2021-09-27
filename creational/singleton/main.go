package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single
func getInstance() *single {
	if singleInstance != nil {
		fmt.Println("Single instance already created.")
		return singleInstance
	}

	lock.Lock()
	defer lock.Unlock()
	if singleInstance != nil {
		fmt.Println("Single instance already created.")
		return singleInstance
	} else {
		fmt.Println("Creating single instance now.")
		singleInstance = &single{}
		return singleInstance
	}
}

func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}