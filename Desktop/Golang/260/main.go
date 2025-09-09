package main

import (
	"fmt"
	"myproject/input"
)

func findMin() (int,int) {
	values := input.InputNumber()
	min := values[0]
	index := 0
	for i,val := range values {
		if val <= min {
			min = val
			index = i
		}
	}
	return min,index
}

func main(){
	min,index := findMin()
	fmt.Printf("Min : %v\nIndex: %v",min,index)
}