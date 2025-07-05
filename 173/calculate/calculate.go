package calculate

import "fmt"

func CalculateValue(N int, A, B float64) float64 {
	if N < 1 || A > B {
		fmt.Println("There is error:")
		return 0 
	}
	var H float64
	H = (B-A) / float64(N)
	return H
}