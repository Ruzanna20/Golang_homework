package main2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputArray(n int) []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text:")
	str,err := reader.ReadString('\n')
	CheckError(err)

	var symbols []string
	str = strings.TrimSpace(str)
	for _,symbol := range str {
		symbols = append(symbols,string(symbol))
	}
	if len(symbols) > n || len(symbols) < n {
		panic("Overflow")
	} 
	return symbols
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}