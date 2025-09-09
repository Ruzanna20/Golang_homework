package calculate

func CalculateFact(N int) int {
	if N == 1 {
		return N
	} 
	return N * CalculateFact(N-1)

	// var mul int = 1
	// for i := 1; i <= N ; i++{
	// 	mul *= i
	// }
	// return mul
} 