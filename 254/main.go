package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var min int = values[0]
	var max int = values[0]
	for _,val := range values {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	artadryal := min * max
	fmt.Println("Artadryal:",artadryal)
}