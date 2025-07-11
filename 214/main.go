package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var sum int = 0
	var count int
	for _,val := range values {
		if val < 0 {
			sum += val
			count++ 
		}
	}
	avg := sum / count
	fmt.Println("Avg:",avg)

}