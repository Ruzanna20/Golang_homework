package main

import (
	"fmt"
	"myproject/main1"
)

func ChangeIndex(arr []int) []int {
	for i := 2; i < len(arr); i += 2 {
		current := arr[i]
		arr[i] = arr[i-1]
		arr[i-1] = current 
	}
	return arr
}

func main() {
	X := main1.InputArray(5)
	Y := ChangeIndex(X)

	fmt.Println("Y:", Y)
}