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

	values := make([]int,len(lines))
	if len(values) > 4{
		fmt.Println("Petq chi sharunakel")
		return 
	}
	var a bool = false	
	for i,line := range lines {
		values[i],_ = strconv.Atoi(line)
		fmt.Println(values[i])
		
	if values[0] == values[1]+values[2]+values[3] || values[1] == values[0]+values[2]+values[3] || values[2] == values[0]+values[1]+values[3] || values[3] == values[0]+values[1]+values[2] {
			a = true
	}
	}
	if a == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}