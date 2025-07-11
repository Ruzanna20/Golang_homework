package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var m int = 2
	var mul int = 1
	for _,val := range values {
		if val % m == 0{
			mul *= val
		}	
	}
	fmt.Println("Artadryal:",mul)
}