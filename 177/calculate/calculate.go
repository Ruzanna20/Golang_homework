package calculate


func Calculate(N int) ([]float64) {
	x := make([]float64, N+1)
	x[1] = 1
	x[2] = 2
	for k := 3 ; k <= N ; k++ {
		x[k] = (x[k-2] + 2*x[k-1]) / 3
	}
	return x
}