package main

import (
	"fmt"
	"myproject/main1"
)

func MinNum(arr []int) int {
	min := arr[0]

	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func newArray(arr []int) []int {
	var newArray []int
	min := MinNum(arr)

	for i,val := range arr {
		if val == min {
			newArray = append(newArray, i)
		}
	}
	return newArray
}

func main() {
	X := main1.InputArray(6)
	Y := newArray(X)

	fmt.Println("Y: ", Y)
}
