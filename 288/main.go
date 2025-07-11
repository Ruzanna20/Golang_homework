package main

import (
	"fmt"
	"myproject/main1"
)

func CheckIndex(arr []int) []int {
	var Y []int
	for i, val := range arr {
		if i%2 == 1 {
			Y = append(Y, val)
		}
	}
	return Y
}

func main() {
	X := main1.InputArray(5)
	Y := CheckIndex(X)

	fmt.Println("Y:",Y)
}