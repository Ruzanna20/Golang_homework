package divide

import "fmt"

func SumNumber(str string, n int) int {
	numbers := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
	var sum int = 0
	values := make([]int,0,100)
	for _,val := range numbers {
		if n % val == 2 {
			sum += val
			values = append(values,val)
		} 
	}
	fmt.Println("Values:",values)
	return sum
}