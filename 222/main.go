package main

import (
	"fmt"
	"inputproject/input"
)


func main() {
	values := input.InputNumber()

	var c int = 5
	var d int = 10
	var mul int = 1
	for _,val := range values {
		if val >= c && val < d {
			mul *= val
		}
	}
	fmt.Println(mul)
}