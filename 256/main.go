package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var min int = values[0]
	var index int 
	for i,val := range values {
		if val < min {
			min = val
			index = i
		}
		 
	}
	sum := index + min
	fmt.Println("Sum:",sum)
}