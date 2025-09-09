package main

import ("fmt"
"project12/calculate")

func main() {
	h,k := calculate.CalculateValue(4,2.0,10.0)
	fmt.Printf("H = %v, k = %v ",h,k)
}