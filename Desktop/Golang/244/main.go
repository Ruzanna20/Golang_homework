package main

import ("myproject/input"
"fmt")

func main() {
	values := input.InputNumber()

	var mul int = 1
	for _,val := range values {
		if val % 5 == 2{
			mul *= val
		}
	}	
	
	fmt.Println("Artadryal:",mul)
}