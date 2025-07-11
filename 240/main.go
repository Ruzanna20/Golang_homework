package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var count int 
	for _,val := range values {
		if val % 7 == 0{
			count++
		}
	}

	fmt.Println("Count:",count)
}