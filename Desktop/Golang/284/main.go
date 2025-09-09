package main

import (
	"fmt"
	"myproject/main1"
)

func InsertArray(arr []int) []int {
	a := 5
	b := 20
	var Y []int
	for _, val := range arr {
		if val >= a && val <= b {
			Y = append(Y, val)
		}
	}
	return Y
}

func main() {
	X := main1.InputArray(10)
	Y := InsertArray(X)

	fmt.Println("Y:", Y)
}