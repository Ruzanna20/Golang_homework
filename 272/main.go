package main

import (
	"fmt"
	"myproject/main2"
)

func CheckSymbols(arr []string, N int) bool {
	t := false
	count := 0
	for _, symbol := range arr {
		if symbol == "b" {
			count++
		}
	}
	if count >=  N/2 {
		t = true
	}
	return t
}

func main() {
	symbols := main2.InputArray(10)

	t := CheckSymbols(symbols,6)
	fmt.Println("T :",t)
}