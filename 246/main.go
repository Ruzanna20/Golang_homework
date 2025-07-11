package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var sum int
	var count int
	for i,val := range values {
		for n := 1 ; n * n <= i ; n++ {
			if n * n == i {
				sum += val
				count ++
			}
		}
	}

	if count != 0 {
		avg := sum / count 
		fmt.Println("Avg:",avg)
	} else {
		fmt.Println("Tivy 0-i chi kareli bajanel:")
	}
}