package main

import (
	"fmt"
	"myproject/main1"
)

func findLessNumThanb(arr []int, b int) []int {
	var result []int

	for _, val := range arr {
		if val < b {
			result = append(result, val)
		}
	}
	return result
}

func newArray(arr []int) []int {
	var newArr []int
	smallNumbers := findLessNumThanb(arr, 5)

	if len(smallNumbers) > 0 {
		for _, val := range arr {
			if val > 0 {
				newArr = append(newArr, val)
			}
		}
	} else {
		for _, val := range arr {
			if val < 0 {
				newArr = append(newArr, val)
			}
		}
	}
	return newArr
}

func main() {
	X := main1.InputArray(5)
	Y := newArray(X)

	fmt.Println("Y: ", Y)
}
