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
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	lines := strings.Fields(str)
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		if val >= 1000 && val <= 9999 {
			fmt.Println(val)

			m := val % 10
			t := (val % 100) / 10
			haryur := (val / 100) % 10
			if m > t {
				result := m * haryur 
				fmt.Println(result)
			}else {
				fmt.Println(1)
			}
		} else {
			fmt.Println("Qaranish tiv che:", val)
		}
	}
}