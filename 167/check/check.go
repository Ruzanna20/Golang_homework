package check

import "math"

func PowerNumber(x float64, n int) float64{
	var value float64 = 1.0
	for i := 0; i <= n ; i++ {
		value *= x 
	}
	return value
}

func CheckValue(x float64) bool {
	var y bool = false
	for n := 1; n <= 30; n++ {
		value := PowerNumber(x,n)
		if math.Sin(value) < 0 {
			y = true
		}
	}
	return y
}