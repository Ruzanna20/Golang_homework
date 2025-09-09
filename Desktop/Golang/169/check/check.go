package check


func CheckNumbers(x int, y int) int {
	var z int = 5
	var count int = 0
	for i := 1; i <= (x + y); i++ {
		if x >= 1 && y >= 1 && (x+y)%i == 0 {
			count += 1
		}
		if count > 2 {
			z = 6
		}
	}
	return z
}