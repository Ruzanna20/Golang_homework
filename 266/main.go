package main

import (
	"fmt"
	"myproject/main1"
)

func main() {
	X := main1.InputArray(5)
	Y := main1.InputArray(10)

	sum1 := 0
	for _,val := range X {
		if val % 2 != 0{
			sum1 += val
		}
	}

	sum2 := 0
	for _,val := range Y {
		if val % 2 == 0{
			sum2 += val
		}
	}

	sub := sum1 - sum2
	fmt.Println("Sub: ",sub)
}
