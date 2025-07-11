package main

import (
	"fmt"
	"myproject/main1"
)

func CountNum(arr []int,k int) int {
	count := 0
	for _, val := range arr {
		if val%k == 0 {
			count++
		}
	}
	return count
}

func main() {
	X := main1.InputArray(5)
	Y := main1.InputArray(10)

	count1 := CountNum(X,2)
	count2 := CountNum(Y,2)
	sumcount := count1 + count2
	fmt.Println("Count: ",sumcount)
}