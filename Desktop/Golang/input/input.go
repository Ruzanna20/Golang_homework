package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputNumber() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter numbers: ")
	str,err := reader.ReadString('\n')
	checkError(err)
	str = strings.TrimSpace(str)
	lines := strings.Fields(str)

	var values []int
	for _,line := range lines {
		val,err := strconv.Atoi(line)
		checkError(err)
		values = append(values, val)
	}
	return values
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
