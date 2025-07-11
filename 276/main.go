package main

import (
	"fmt"
	"myproject/main2"
)

func CheckSymbols(arr []string) bool {
	t := false

	for _, symbol := range arr {
		if symbol == "r" {
			t = true
		}
	}
	return t
}

func main() {
	symbols := main2.InputArray(5)
	
	result := CheckSymbols(symbols)
	fmt.Println(result)
}