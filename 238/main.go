package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var sum int = 0
	var count int = 0
	for _,val := range values {
		if val % 3 == 0 {
			sum += val
			count++
		}
	}

	if count != 0 {
		avg := sum / count
		fmt.Println("Avg:",avg)
	} else {
		fmt.Println("3-i bazmapatik chka:")
	}
}