package main

import (
	"fmt"
	"myproject/input"
	"math"
)

func main() {
	values := input.InputNumber()

	var k int = 2
	var sum int = 0
	var count int 
	for _,val :=  range values {
		if val % k == 0 {
			sum += val* val
			count ++ 
		}
	}

	mijqar := int(math.Sqrt(float64(sum)/float64(count)))
	fmt.Println("Mijin qarakusayin:",mijqar)
}