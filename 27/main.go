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
	fmt.Println("Enter the text: ")
	str,_ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	lines := strings.Fields(str)

	values := make([]int,len(lines)) //chatic em nayel,vonc stringy darcnem slice
	for i,line := range lines{
		values[i],_  = strconv.Atoi(line)
		fmt.Println(values[i]) }

	var a bool = true
	d := values[1] - values[0] 
	for i := 1; i < len(values);i++ {
			if values[i] - values[i-1] != d{
				a = false
			} 	
			}
	if a == true {
		fmt.Println("True")
	} else{
		fmt.Println("False")
	}
}
