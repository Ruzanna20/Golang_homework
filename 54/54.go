package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number:")
	str,_ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	lines := strings.Fields(str)
	for _,line := range lines {
		val,_ := strconv.Atoi(line)
		if val >= 100 && val <= 999{
			fmt.Println(val)
		
			m := val % 10
			t := (val / 10) % 10
			h := val / 100 
			max := m
			if t > max {
				max = t
			}
			if h > max {
				max = h
			}
			fmt.Println(max)
		} else{
			fmt.Println("Eranish tiv che:",val)
		}

}
}
