package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	//Insert values
	m.Store("Anna",1)
	m.Store("Ani",2)

	// Get value
	value,okey := m.Load("Anna")
	if okey {
		fmt.Println("Key found:",value)
	} 

	// Get or store value
	val,loaded := m.LoadOrStore("Sona",3)
	fmt.Println(val,loaded) 

	// Delete key
	m.Delete("Anna")

	// Iteration
	m.Range(func(key, value any) bool {
		fmt.Println(key,value)
		return true
	})
}

