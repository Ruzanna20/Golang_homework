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

func MinNum(arr []int) int {
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func avgOfMaxAndMin(arr []int) int {
	max := MaxNum(arr)
	min := MinNum(arr)
	avg := (max + min) / 2
	return avg
}

func newArray(arr []int) []int {
	var newArray []int
	avg := avgOfMaxAndMin(arr)
	for _, val := range arr {
		if val > avg {
			newArray = append(newArray, val)
		}
	}
	return newArray
}

func main() {
	X := main1.InputArray(5)
	Y := newArray(X)

	fmt.Println("Y: ", Y)
}
