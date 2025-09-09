package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the text: ")
	str,_ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	lines := strings.Fields(str)
	var a bool = false
	for _,line := range lines{
		val,_ := strconv.Atoi(line)
		fmt.Println(val) 

		if val == 1 {
			a = true
			}
		}
		if a == true {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
}
