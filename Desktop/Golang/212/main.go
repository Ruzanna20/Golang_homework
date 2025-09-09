package main

import (
	"fmt"
	"myproject/input"
	"math"
)

func main() {
	values := input.InputNumber()
	var sum int = 0
	var count int
	for _,val := range values {
		sum += val * val
		count++
	}

	mijqar := math.Sqrt(float64(sum) / float64(count))
	fmt.Println("Mijin qarakusayin:",mijqar)
}





