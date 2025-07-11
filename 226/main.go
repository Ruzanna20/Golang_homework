package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var count int
	var k int = 5
	for _,val := range values {
		if val > 0 {
			if val < k {
				count++
			}
		} else if val < 0 {
			if -val < k {
				count++
			}
		}
	}

	fmt.Println("Count:",count)
}