package main

import (
	"fmt"
	"inputproject/input"
)

func main() {
	values := input.InputNumber()

	var sum int = 0
	for i,val := range values {
		if i % 2 == 1 {
			if val >= 0 {
				sum += val
			} else {
				sum += -val
			}
		}
	}
	fmt.Println("Sum:",sum)
}