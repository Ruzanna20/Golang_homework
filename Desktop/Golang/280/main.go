package main

import (
	"fmt"
	"myproject/main2"
)

func NewArray(arr []string) []string {
	var newarray []string

	for _, symbol := range arr {
		newarray = append(newarray, symbol)
		if symbol == "f" {
			newarray = append(newarray, symbol)
		}
	}
	return newarray
}

func main() {
	symbols := main2.InputArray(5)
	result := NewArray(symbols)
	fmt.Println(result)

}