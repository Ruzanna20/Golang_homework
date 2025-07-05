package main

import "fmt"

func calculate(N int, K int) (int, int) {
	var count int = 0
	for N >= K {
		N = N - K
		count++
	}
	return N, count
}

func main() {
	mnacord,qanord:= calculate(8, 4)
	fmt.Printf("Mnacord : %v\nQanord : %v",mnacord,qanord)
}