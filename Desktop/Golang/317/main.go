package main

import (
	"fmt"
	"myproject/main1"
)

func MaxNum(arr []int) int {
	max := arr[0]
	for _, val := range arr {
		if val%2 == 0 {
			if val > max {
				max = val
			}
		}
	}
	return max
}
func newArray(arr []int) []int {
	var newArray []int

	for i, val := range arr {
		if i%2 != 0 {
			val += MaxNum(arr)
			newArray = append(newArray, val)
		}
	}
	return newArray
}

func main() {
	X := main1.InputArray(10)
	Y := newArray(X)

	fmt.Println("Y: ", Y)
}
