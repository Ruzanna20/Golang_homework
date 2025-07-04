package check

import "fmt"

func SumNumber() int {
	numbers := make([]int, 0, 90)
	for i := 10; i <= 99; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
	var sum int = 0
	const n int = 3
	values := make([]int,0,90)
	for _,val := range numbers {
		if val % n == 0 {
			sum += val
			values = append(values, val)
		}
	}
	fmt.Println("Values:",values)
	return sum
}