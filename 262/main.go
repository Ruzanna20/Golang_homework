package main

import (
	"fmt"
	"inputproject/package1"
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
	A := main1.InputArray(5)
	B := main1.InputArray(10)

	mijqarA := mijqar(A)
	mijqarB := mijqar(B)
	sum := mijqarA + mijqarB
	fmt.Println("Sum:",sum)
}