package main

import "fmt"

func main() {
	koxmer := make([]int, 0, 3)
	koxmer = append(koxmer, 1, 1, 2)
	var y int
	if koxmer[0]+koxmer[1] > koxmer[2] && koxmer[0]+koxmer[2] > koxmer[1] && koxmer[1]+koxmer[2] > koxmer[0] {
		y = 1
	} else {
		y = 2
	}
	fmt.Println(y)
}
