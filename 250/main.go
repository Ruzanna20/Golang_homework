package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var mul int = 1
	for i,val := range values {
		if i * val % 3 == 2{
			mul *= val
		}
	}
	
	fmt.Println("Artadryal:",mul)
}