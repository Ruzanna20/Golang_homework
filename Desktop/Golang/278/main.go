package main

import (
	"fmt"
	"myproject/main2"
)

func NewArray(arr []string) []string {
	var newarray []string

	for i, symbol := range arr {
		if i%2 != 0 {
			newarray = append(newarray, symbol)
		}
	}
	return newarray
}

func main() {
	symbols := main2.InputArray(5)
	newarr := NewArray(symbols)
	fmt.Println("New Array:",newarr)
}