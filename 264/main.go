package main

import (
	"fmt"
	"myproject/main1"
)

func CountNum(arr []int) int {
	count := 0
	for _,val := range arr {
		if val < 0{
			count++
		}
	}
	return count	
}

func main() {
	X := main1.InputArray(5)
	Y := main1.InputArray(10)

	count1 := CountNum(X)
	count2 := CountNum(Y)
	sum := count1 + count2
	fmt.Println("Bacasakan tveri qanaky: ",sum)
}