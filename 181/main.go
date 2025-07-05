package main

import "fmt"

func calculate(N int) int {
	var val int = 1
	var K int 
	for val < N {
		val *= 2
		K++ 
	}
	return K
}

func main() {
	result := calculate(16)
	fmt.Println("K : ",result)
}