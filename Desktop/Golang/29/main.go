package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i,line := range lines {
			values[i],_ = strconv.Atoi(line)
			fmt.Println(values[i])
		}

	sort.Ints(values) // https://pkg.go.dev/sort#Sort
	fmt.Println(values)
	
}