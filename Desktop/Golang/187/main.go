package main

import "fmt"

func checkNumber(N int) bool {
	var count int
	var y bool = false
	for i := 1; i <= N; i++ {
		if N % i == 0 {
			count++
		}
	}
	if count == 2 {
		y = true
	}
	return y
}

func main() {
	result := checkNumber(13)
	fmt.Println(result)
}