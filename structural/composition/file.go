package main

import "fmt"

// Component interface
type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keword %s in file %s", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}


