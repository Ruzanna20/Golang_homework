package main

import (
	"fmt"
	"myproject/main1"
)

func MaxNum(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func NewArr(arr []int) []int {
	var newarr []int
	max := MaxNum(arr)
	for _, val := range arr {
		if val > 0 {
			val += max
			newarr = append(newarr, val)
		}
	}
	return newarr
}

func main() {
	X := main1.InputArray(5)
	Y := NewArr(X)

	fmt.Println("Y:", Y)
}