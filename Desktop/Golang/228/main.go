package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var k int = 2
	var sum int = 0
	for i,val := range values {
		if i % k == 0 {
			sum += val
		}
	}

	fmt.Println("Sum:",sum)
}