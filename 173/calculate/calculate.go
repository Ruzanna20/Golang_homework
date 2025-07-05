package calculate

import "fmt"

func CalculateValue(N int, A, B float64) (float64,[]float64) {
	if N < 1 || A > B {
		fmt.Println("There is error:")
		return 0,nil
	}
	var H float64
	H = (B-A) / float64(N)
	var koordinat []float64
	for i := 0 ; i <= N ; i++{
		k := A + (float64(i) * H)
		koordinat = append(koordinat,k )
	}
	return H,koordinat

}