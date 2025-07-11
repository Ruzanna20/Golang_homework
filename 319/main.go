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

func newArray(arr []int) []int {
	var newArray []int
	max := MaxNum(arr)

	for i := 0; i < len(arr); i++ {
		if i%3 == 0 {
			arr[i] = 0
			newArray = append(newArray, arr[i])
		} else {
			arr[i] += max
			newArray = append(newArray, arr[i])
		}
	}
	return newArray
}

func main() {
	X := main1.InputArray(10)
	Y := newArray(X)

	fmt.Println("Y: ", Y)
}
