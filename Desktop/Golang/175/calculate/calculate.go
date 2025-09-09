package calculate 

func Calculate(N int) []float64 {
	x := make([]float64,N+1)
	x[0] = 1
	for k := 1 ; k <= N ; k ++ {
			x[k] = (x[k-1] + 1) / float64(k)
		}
	return x
}