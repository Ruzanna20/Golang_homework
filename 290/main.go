package main

import (
	"fmt"
	"myproject/main1"
)

func CheckValue(arr []int) []int {
	var Y []int
	for _, val := range arr {
		if val%6 == 1 {
			Y = append(Y, val)
		}
	}
	return Y
}

func main() {
	X := main1.InputArray(5)
	Y := CheckValue(X)

	fmt.Println("Y:",Y)
}