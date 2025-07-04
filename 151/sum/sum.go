package sum

import "fmt"

func SumNumber(str string, n int) int {
	numbers := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
	var sum int = 0
	for _,val := range numbers {
		if n % val == 0 {
			sum += val
		}
	}
	return sum
}