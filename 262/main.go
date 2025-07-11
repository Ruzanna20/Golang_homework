package main

import (
	"fmt"
	"myproject/main1"
	"math"
)

func mijqar(arr []int) int {
	sum := 0
	count := 0
	
	for _,val := range arr {
		sum += val * val
		count++
	}
	return int(math.Sqrt(float64(sum) / float64(count)))
}

func main() {
	X := main1.InputArray(5)
	Y := main1.InputArray(10)

	mijqarA := mijqar(X)
	mijqarB := mijqar(Y)
	sum := mijqarA + mijqarB
	fmt.Println("Sum:",sum)
}