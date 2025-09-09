package main

import (
	"fmt"
	"myproject/main1"
)

func countDuplicateInArr(arr []int) map[int]int {
	count := make(map[int]int) 

	for _, val := range arr {
		count[val]++

	}
	return count
}

func newArray(arr []int) []int {
	var newArray []int
	count := countDuplicateInArr(arr)

	for _, val := range arr {
		if count[val] == 1 {
			newArray = append(newArray, val)
		}
	}
	return newArray
}

func main() {
	X := main1.InputArray(6)
	Y := newArray(X)

	fmt.Println("Y: ", Y)
}
