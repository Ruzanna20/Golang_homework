package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var min int = values[0]
	for _,val := range values  {
		if val < min {
			min = val
		}
	}

	fmt.Println("Min։", min)
}