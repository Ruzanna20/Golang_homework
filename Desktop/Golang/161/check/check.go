package check

func FindNumber() int {
	values := make([]int,0,9000)
	for i := 1000;i<=9999;i++{
		for a := 1 ; a <= 9999 ; a++{
			if i * 26 == a * a  {
				values = append(values,i)
			}
		}	
	}
	var min int = values[0]
	for _,value := range values {
		if value < min {
			min = value
		}
	}
	return min
} 