package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var sum int 
	var k int = 3
	for i,val := range values {
		if ((i + val) * (i + val)) % k == 0 {
			sum += val
		}
	} 
	fmt.Println("Sum:",sum)
	 
}