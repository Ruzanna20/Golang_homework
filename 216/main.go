package main

import (
	"fmt"
	"inputproject/input"
)

func main() {
	values := input.InputNumber()

	var mul int = 1
	for i,val := range values {
		if i % 2 == 0{
			mul *= val
		}
	}

	fmt.Println(mul)
}