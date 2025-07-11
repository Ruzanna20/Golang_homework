package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var sum int 
	var count int
	for _,val := range values {
		if val % 2 == 1 {
			sum += val
			count++
		}
	}

	if count != 0{
		avg := sum / count
		fmt.Println("Avg:",avg)
	} else {
		fmt.Println("Kent tiv chka:")
	}
}