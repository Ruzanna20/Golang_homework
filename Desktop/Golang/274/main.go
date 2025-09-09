package main

import (
	"fmt"
	"myproject/main2"
)

func CalculateAvg(arr []string) int {
	sum := 0
	count := 0
	var index []int
	for i, symbol := range arr {
		if symbol > "h" {
			index = append(index, i)
		}
	}
	fmt.Println(index)
	for _,val := range index {
		sum += val
		count++
	}

	avg := sum / count 
	return avg
}

func main() {
	symbols := main2.InputArray(5)
	avg := CalculateAvg(symbols)
	fmt.Println("Avg:",avg)
}