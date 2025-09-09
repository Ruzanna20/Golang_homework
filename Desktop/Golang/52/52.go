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
	fmt.Println("Enter the number: ")
	str,_ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	lines := strings.Fields(str)
	for _,line := range lines {
		val,_ := strconv.Atoi(line)
		if val >= 100 && val <= 999{
			fmt.Println(val)
			
			m := val % 10
			t1 := (val / 10) % 10
			h := val / 100 
			var t bool = false
			if m == t1 || m ==h || t1 == h {
				t = true
			}
			if t {
				fmt.Println("True")
			} else {
				fmt.Println("False")
			}
		} else{
			fmt.Println("Eranish tiv che:",val)
		}
}
}

