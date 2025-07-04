package check

func CheckNumber(n int) bool {
	var t bool = false
	var values []int
	for m := 3;m <= n ; m *=3 {
		values = append(values,m)
	}
	for _,val := range values {
		if n == val {
			t = true
		}
	}
	return t
}