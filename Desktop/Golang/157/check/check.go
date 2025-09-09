package check

import "fmt"

func SumNumber() int {
	numbers := make([]int, 0, 900)
	for i := 100; i <= 999; i++ {
		numbers = append(numbers, i)
	}
	var sum int = 0
	values := make([]int,0,900)
	for _,val := range numbers {
		if val % 5 != 0 {
			sum += val
			values = append(values, val)
		}
	}
	fmt.Println("Values",values)
	return sum
}