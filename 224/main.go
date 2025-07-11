package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var k int = 5
	var sum int = 0
	for _,val := range values {
		if val > 0 {
			if val < k {
				sum += val * val * val
			}
		} else if val < 0 {
			if -val < k {
				sum += val * val * val
			}
		}
	} 

	fmt.Println("Sum:",sum)
}