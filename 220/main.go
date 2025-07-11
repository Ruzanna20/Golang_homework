package main

import (
	"fmt"
	"myproject/input"
)

func main() {
	values := input.InputNumber()

	var drakan int
	var bacasakan int
	for _,val := range values {
		if val > 0 {
			drakan ++
		} else if val < 0 {
			bacasakan++ 
		} else {
			fmt.Println("0 nshan chuni:")
		}
	}
	fmt.Printf("Drakan: %v\nBacasakan: %v",drakan,bacasakan)
}