package main

import (
	"fmt"
	"inputproject/package1"
)

func CountNum(arr []int) int {
	count := 0
	for _, val := range arr {
		if val < 0 {
			count++
		}
	}
	return count
}

func main() {
	A := main1.InputArray(5)
	B := main1.InputArray(10)

	count1 := CountNum(A)
	count2 := CountNum(B)
	sumcount := count1 + count2
	fmt.Println("Bacasakan tarreri qanaky:",sumcount)
}