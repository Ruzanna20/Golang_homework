package main

import "fmt"

func calculate(avand float64,P float64) (float64, int) {
	var count int
	for avand <= 100000 {
		avand += avand * P
		count++
	}
	return avand, count
}

func main() {
	avand, amis := calculate(30000.0,0.15)
	fmt.Printf("Avand : %v\nAmisneri qanak: %v",avand,amis)
}