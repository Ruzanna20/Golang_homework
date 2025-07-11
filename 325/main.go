package main

import (
	"fmt"
	"myproject/main1"
)

func MaxNum(arr []int) (int, int) {
	max := arr[0]
	maxIndex := 0
	for i, val := range arr {
		if val > max {
			max = val
			maxIndex = i
		}
	}
	return max, maxIndex
}

func newArray(arr []int) []int {
	var newArray []int
	// max, maxIndex := MaxNum(arr)
	// arr[maxIndex] = arr[0]
	// arr[0] = max
	_, maxIndex := MaxNum(arr)
	arr[0], arr[maxIndex] = arr[maxIndex], arr[0] // chatic em harcrel es tarberaky

	for i := 0; i < len(arr); i++ {
		newArray = append(newArray, arr[i])
	}
	return newArray
}

func main() {
	X := main1.InputArray(10)
	Y := newArray(X)

	fmt.Println("Y: ", Y)
}
