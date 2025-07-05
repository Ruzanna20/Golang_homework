package main

import "fmt"

func calculate(N int) int {
	if N <= 1 {
		fmt.Println("There is error:")
		return 0
	}
	var val int = 1
	var K int 
	for val <= N {
		val *= 3
		K++
	}
	var min int
	if val > N {
		min = K
	}
	return min
}

func main() {
	result := calculate(9)
	fmt.Println("K = ",result)
}