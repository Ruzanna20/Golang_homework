package main

import (
	"fmt"
	"myproject/main1"
)

func MaxNum(arr []int) (int, int) {
	max := arr[3]
	maxIndex := 3
	for i := 3; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}
	return maxIndex, max
}

func newArray(arr []int) []int {
	var newArray []int
	maxIndex, max := MaxNum(arr)
	for i, val := range arr {
		if i%2 == 0 && i < maxIndex {
			val += max
			newArray = append(newArray, val)
		}
	}
	return newArray
}

func main() {
	X := main1.InputArray(7)
	Y := newArray(X)

	fmt.Println("Y: ", Y)
}
