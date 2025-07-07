package main

import (
	"fmt"
	"inputproject/main1"
)

func ChangeIndex(arr []int) []int {
	for i := 1; i < len(arr)-1; i += 1 {
		current := arr[i+1]
		arr[i+1] = arr[i]
		arr[i] = current 
	}
	return arr
}

func main() {
	X := main1.InputArray(5)
	Y := ChangeIndex(X)

	fmt.Println("Y:", Y)
}