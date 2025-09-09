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
			hazar := val / 1000
			if m + haryur != 0 {
			if val < 5000 {
				result := val / (m + haryur)
				fmt.Println(result)
			} else {
				result := val / (hazar + t)
				fmt.Println(result)
			}
			} else {
				fmt.Println("m + haryur = 0")
			}

		} else {
			fmt.Println("Qaranish tiv che:", val)
		}
	}
}