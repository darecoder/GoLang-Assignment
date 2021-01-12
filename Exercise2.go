package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inp := os.Args[1]
	num, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println("Not a number:", inp)
		os.Exit(1)
	}
	if num%2 == 0 {
		fmt.Println("even number:", num)
	} else {
		fmt.Println("odd number:", num)
	}
}
