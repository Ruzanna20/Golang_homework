package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var max int = values[0]
	var index int
	for i,val := range values {
		if val >= max {
			max = val
			index = i
		}
	}
	fmt.Printf("Max: %v\nIndex: %v",max,index)
}