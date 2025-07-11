package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var count int
	var mul int = 1
	for _,val := range values {
		if val % 2 != 0 {
			mul *= val
			count++
		}
	}

	fmt.Printf("Count: %v\nArtadryal: %v",count,mul)
}